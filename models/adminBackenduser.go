package models

import "github.com/astaxie/beego/orm"

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
