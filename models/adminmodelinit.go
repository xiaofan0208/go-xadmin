package models

import "github.com/xiaofan0208/xbase/utils"

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
