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
	var user *Backenduser
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
	return user, nil
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
