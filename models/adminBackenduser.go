package models

import (
	"errors"

	"github.com/astaxie/beego/orm"
	"github.com/xiaofan0208/xbase/utils"
)

//Backenduser 管理用户
type Backenduser struct {
	Id       int64  `orm:"column(id);pk;auto" json:"id"`
	Name     string `orm:"column(name)" json:"name"`
	Email    string `orm:"column(email)" json:"email"`       // 邮箱
	Mobile   string `orm:"column(mobile)" json:"mobile"`     // 手机号
	Avatar   string `orm:"column(avatar)" json:"avatar"`     // 头像
	Password string `orm:"column(password)" json:"-"`        // 密码
	Salt     string `orm:"column(salt)" json:"salt"`         // 加盐
	IsAdmin  bool   `orm:"column(is_admin)" json:"is_admin"` // 是否是管理员
	Active   bool   `orm:"column(active)" json:"active"`     // 是否可以使用，默认是
	Deleted  bool   `orm:"column(deleted)" json:"deleted"`   // 是否删除
	Created  int64  `orm:"column(created)" json:"created"`   // 创建时间
	Updated  int64  `orm:"column(updated)" json:"updated"`   // 更新时间
}

// 更新数据字段
var BackenduserUpdateFields = []string{"Id", "Name", "Email", "Mobile", "Avatar", "IsAdmin", "Active", "Updated"}

//TableName 表名
func (u *Backenduser) TableName() string {
	return "admin_backenduser"
}

//TableUnique 多字段唯一键
func (u *Backenduser) TableUnique() [][]string {
	return [][]string{
		[]string{"Name"},
	}
}

func init() {
	orm.RegisterModel(new(Backenduser))
}

// GetBackenduserByID get
func GetBackenduserByID(id int64) (*Backenduser, error) {
	user := Backenduser{}
	o := orm.NewOrm()
	// 获取 QuerySeter 对象，user 为表名
	qs := o.QueryTable(new(Backenduser))
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

// CheckUserByName get
func CheckUserByName(username string, password string) (*Backenduser, error) {
	var user Backenduser
	var err error
	o := orm.NewOrm()

	cond := orm.NewCondition()
	cond1 := cond.And("Active", true).And("Email", username).Or("Mobile", username)

	qs := o.QueryTable(new(Backenduser))
	qs = qs.SetCond(cond1)

	if err = qs.One(&user); err == nil {
		if user.Deleted == true || user.Active == false {
			return nil, errors.New("用户无效")
		}
		if user.Password != utils.EncodePassword(password, user.Salt) {
			return nil, errors.New("密码错误")
		}
		return &user, nil
	}
	if err == orm.ErrNoRows {
		return nil, errors.New("用户不存在")
	}
	return &user, err
}

// CreateBackenduser 创建管理员
func CreateBackenduser(record *Backenduser) (*Backenduser, error) {
	o := orm.NewOrm()
	record.Deleted = false
	record.Active = true
	record.Created = utils.NowMillis()
	record.Updated = record.Created

	id, err := o.Insert(record)
	if err == nil {
		record.Id = id
		return record, nil
	}
	return nil, err
}

// GetBackenduserByParam 根据参数查找
func GetBackenduserByParam(params map[string]interface{}, offset int64, limit int64) ([]*Backenduser, int64, error) {
	o := orm.NewOrm()
	users := []*Backenduser{}
	qs := o.QueryTable(new(Backenduser))
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

	_, err := qs.All(&users)
	if err == orm.ErrNoRows {
		return users, 0, nil
	}
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

//GetAllBackenduserList 获取所有列表
func GetAllBackenduserList(params map[string]interface{}, exclude map[string]interface{}, orders []string, limit uint8, offset uint8) ([]*Backenduser, int64, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Backenduser))
	var records []*Backenduser
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

// UpdateBackenduserByField 更新
func UpdateBackenduserByField(record *Backenduser, fields ...string) (int64, error) {
	o := orm.NewOrm()
	var (
		num int64
		err error
	)

	if err = o.Read(&Backenduser{Id: record.Id}); err == nil {
		record.Updated = utils.NowMillis()
		if num, err = o.Update(record, fields...); err == nil {
		}
	} else if err == orm.ErrNoRows {
		// 未查询到
		return num, nil
	}
	return num, err
}

// DeleteBackenduser 物理删除
func DeleteBackenduser(id int64) (int64, error) {
	o := orm.NewOrm()
	var (
		num int64
		err error
	)
	if num, err = o.Delete(&Backenduser{Id: id}); err == nil {
		return num, err
	}
	return num, nil
}

// DeleteBackenduserByID 逻辑删除
func DeleteBackenduserByID(id int64) error {
	o := orm.NewOrm()
	record := Backenduser{Id: id}
	record.Deleted = true
	record.Updated = utils.NowMillis()

	var err error
	if _, err = o.Update(&record, "Deleted", "Updated"); err == nil {
		return nil
	}
	return err
}
