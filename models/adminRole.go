package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/xiaofan0208/go-xbase/utils"
)

//AdminRbacRole 角色
type AdminRbacRole struct {
	Id      int64  `orm:"column(id);pk;auto" json:"id"`
	Pid     int64  `orm:"column(pid)" json:"pid"` // 父资源ID
	Name    string `orm:"column(name)" json:"name"`
	Desc    string `orm:"column(desc)" json:"desc"`       // 描述
	Type    int64  `orm:"column(type)" json:"type"`       // 角色类型 1-基础角色
	Level   int64  `orm:"column(level)" json:"level"`     // 层级
	Status  uint8  `orm:"column(status)" json:"status"`   // 状态：1-启用 2-禁用
	Deleted bool   `orm:"column(deleted)" json:"deleted"` // 是否删除
	Created int64  `orm:"column(created)" json:"created"` // 创建时间
	Updated int64  `orm:"column(updated)" json:"updated"` // 更新时间
}

// 更新数据字段
var AdminRbacRoleUpdateFields = []string{"Id", "Name", "Status", "Updated"}

//TableName 表名
func (u *AdminRbacRole) TableName() string {
	return "admin_rbac_role"
}

func init() {
	orm.RegisterModel(new(AdminRbacRole))
}

//GetAllAdminRbacRoleList 获取所有列表
func GetAllAdminRbacRoleList(params map[string]interface{}, exclude map[string]interface{}, orders []string, limit uint8, offset uint8) ([]*AdminRbacRole, int64, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(AdminRbacRole))
	var records []*AdminRbacRole
	var err error
	var count int64
	// 过滤
	for k, v := range exclude {
		qs = qs.Exclude(k, v)
	}
	// 查询参数
	for k, v := range params {
		qs = qs.Filter(k, v)
	}

	// 排序
	qs = qs.OrderBy(orders...)

	// 分页
	if limit == 0 {
		limit = 20
	}
	qs = qs.Limit(limit, offset)
	qs = qs.Distinct()

	count, err = qs.Count()
	if err != nil {
		return records, count, err
	}

	if _, err = qs.All(&records); err == nil {

	}
	// 没有找到
	if err == orm.ErrNoRows {
		return records, count, nil
	}
	return records, count, err
}

// CreateAdminRbacRole 创建管理员
func CreateAdminRbacRole(record *AdminRbacRole) (*AdminRbacRole, error) {
	o := orm.NewOrm()
	record.Deleted = false
	record.Created = utils.NowMillis()
	record.Updated = record.Created

	id, err := o.Insert(record)
	if err == nil {
		record.Id = id
		return record, nil
	}
	return nil, err
}

// GetAdminRbacRoleByID get
func GetAdminRbacRoleByID(id int64) (*AdminRbacRole, error) {
	user := AdminRbacRole{}
	o := orm.NewOrm()
	// 获取 QuerySeter 对象，user 为表名
	qs := o.QueryTable(new(AdminRbacRole))
	qs = qs.Filter("Id", id)
	err := qs.One(&user)
	if err == orm.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateAdminRbacRolerByField 更新
func UpdateAdminRbacRolerByField(record *AdminRbacRole, fields ...string) (int64, error) {
	o := orm.NewOrm()
	var (
		num int64
		err error
	)

	if err = o.Read(&AdminRbacRole{Id: record.Id}); err == nil {
		record.Updated = utils.NowMillis()
		if num, err = o.Update(record, fields...); err == nil {
		}
	} else if err == orm.ErrNoRows {
		// 未查询到
		return num, nil
	}
	return num, err
}

// DeleteAdminRbacRoleByID 逻辑删除
func DeleteAdminRbacRoleByID(id int64) error {
	o := orm.NewOrm()
	record := AdminRbacRole{Id: id}
	record.Deleted = true
	record.Updated = utils.NowMillis()

	var err error
	if _, err = o.Update(&record, "Deleted", "Updated"); err == nil {
		return nil
	}
	return err
}
