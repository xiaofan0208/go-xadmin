package models

import (
	"github.com/xiaofan0208/go-xbase/utils"
)

// InsertAdminUser 插入用户
func InsertAdminUser() {
	user := &Backenduser{
		Name:     "admin",
		Email:    "admin@123.com",
		Password: "123456",
		IsAdmin:  true,
	}
	user.Salt = utils.RandomString(6)
	user.Password = utils.EncodePassword(user.Password, user.Salt)

	CreateBackenduser(user)
}

var menus []*MenuResource

// AddMenus 添加菜单
func AddMenus(menu *MenuResource) {
	if menu == nil {
		return
	}
	menus = append(menus, menu)
}

// InsertMenus 插入菜单
func InsertMenus() {
	InsertMenuResource(menus)
}
