package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/xiaofan0208/go-xbase/utils"
)

// MenuResource 资源
type MenuResource struct {
	Id      int64  `orm:"column(id);pk;auto" json:"id"`
	Title   string `orm:"column(title)" json:"title"`     // 资源标题
	Pid     int64  `orm:"column(pid)" json:"pid"`         // 父资源ID
	Type    int64  `orm:"column(type)" json:"type"`       // 类型 1：panel 2：菜单 3：按钮
	Name    string `orm:"column(name)" json:"name"`       // 名称：唯一标记
	Icon    string `orm:"column(icon)" json:"icon"`       // 图标
	Level   int64  `orm:"column(level)" json:"level"`     // 层级
	Link    string `orm:"column(link)" json:"link"`       // 链接
	UrlFor  string `orm:"column(url_for)" json:"url_for"` // URL
	Desc    string `orm:"column(desc)" json:"desc"`       // 描述
	Order   int64  `orm:"column(order)" json:"order"`     // 排序
	Status  uint8  `orm:"column(status)" json:"status"`   // 状态：1-启用 2-禁用
	Created int64  `orm:"column(created)" json:"created"` // 创建时间
	Updated int64  `orm:"column(updated)" json:"updated"` // 更新时间
	Deleted bool   `orm:"column(deleted)" json:"deleted"` // 是否删除

	Children []*MenuResource `orm:"-" json:"children"`
	Active   bool            `orm:"-" json:"active"`
}

// 更新数据字段
var MenuResourceUpdateFields = []string{"Title", "Email", "Mobile", "Avatar", "IsAdmin", "Active", "Updated"}

//TableName 表名
func (u *MenuResource) TableName() string {
	return "admin_rbac_resource"
}

func init() {
	orm.RegisterModel(new(MenuResource))
}

// InsertMenuResource 插入菜单
func InsertMenuResource(menus []*MenuResource) error {

	// 获取所有菜单
	param := make(map[string]interface{})
	param["Deleted"] = false
	param["Status"] = STATUS_NORMAL
	allmenus, _, err := GetMenuResourceByParam(param, 0, 0)
	if err != nil {
		return err
	}
	allMenusMap := make(map[int64]*MenuResource)
	for _, m := range allmenus {
		allMenusMap[m.Id] = m
	}

	o := orm.NewOrm()
	err = o.Begin()
	for i := 0; i < len(menus); i++ {
		err = insertMenuResource(o, allMenusMap, menus[i], 0)
		if err != nil {
			break
		}
	}

	// 删除所有没有的数据
	for _, m := range allMenusMap {
		m.Deleted = true
		_, err = o.Update(m)
	}

	if err != nil {
		err = o.Rollback()
	} else {
		err = o.Commit()
	}
	return err
}

func insertMenuResource(o orm.Ormer, allMenusMap map[int64]*MenuResource, menu *MenuResource, pid int64) error {
	record := MenuResource{Name: menu.Name}
	record.Title = menu.Title
	record.Type = menu.Type
	record.Name = menu.Name
	record.Icon = menu.Icon
	record.UrlFor = menu.UrlFor
	record.Deleted = false
	record.Created = utils.NowMillis()
	record.Updated = record.Created
	record.Status = STATUS_NORMAL
	record.Pid = pid

	var isExist bool = false
	var err error
	// 查找是否已经存在
	for id, m := range allMenusMap {
		if m.Name == menu.Name {
			isExist = true
			record.Id = id
			delete(allMenusMap, id)
			// 如果已经存在，则更新
			if _, err = o.Update(&record); err == nil {
			} else {
				break
			}
			break
		}
	}
	// 如果不存在，则插入
	if isExist == false {
		var id int64
		if id, err = o.Insert(&record); err == nil {
			record.Id = id
		}
	}
	if err == nil {
		for i := 0; i < len(menu.Children); i++ {
			err = insertMenuResource(o, allMenusMap, menu.Children[i], record.Id)
			if err != nil {
				break
			}
		}
	}
	return err
}

// GetMenuResourceByParam 获取
func GetMenuResourceByParam(params map[string]interface{}, offset int64, limit int64) ([]*MenuResource, int64, error) {
	o := orm.NewOrm()
	records := []*MenuResource{}
	qs := o.QueryTable(new(MenuResource))
	for k, v := range params {
		qs = qs.Filter(k, v)
	}

	total, _ := qs.Count()

	if limit != 0 {
		if offset == 0 {
			qs = qs.Limit(limit)
		} else {
			qs = qs.Limit(limit, offset)
		}
	}

	_, err := qs.All(&records)
	if err == orm.ErrNoRows {
		return records, 0, nil
	}
	if err != nil {
		return nil, 0, err
	}

	return records, total, nil
}
