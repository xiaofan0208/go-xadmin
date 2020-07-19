package routers

import (
	"github.com/xiaofan0208/go-xadmin/controllers"

	"github.com/astaxie/beego"
)

func init() {

}

// Router 后台路由
func Router() {
	beego.Router("/admin/login", &controllers.SignInController{})
	// 所有 /admin/* 路由需要检测是否登录
	beego.InsertFilter("/admin/*", beego.BeforeRouter, LoginFilterFunc)
	// 后台首页
	beego.Router("/admin", &controllers.AdminIndexController{}, "*:Index")
	// 管理员管理
	beego.Router("/admin/backenduser", &controllers.BackenduserController{}, "*:Index")
	beego.Router("/admin/backenduser/list", &controllers.BackenduserController{}, "*:PostList")
	beego.Router("/admin/backenduser/delete", &controllers.BackenduserController{}, "*:DeleteBatch")
	beego.Router("/admin/backenduser/edit/?:id([0-9]+)", &controllers.BackenduserController{}, "*:Edit")
	beego.Router("/admin/backenduser/create", &controllers.BackenduserController{}, "*:Create")

}
