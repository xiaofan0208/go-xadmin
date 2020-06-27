package routers

import (
	"github.com/xiaofan0208/xadmin/controllers"

	"github.com/astaxie/beego"
)

func init() {

}

// Router 后台路由
func Router() {
	beego.Router("/admin/login", &controllers.SignInController{})
	// 所有 /admin/* 路由需要检测是否登录
	beego.InsertFilter("/admin/*", beego.BeforeRouter, LoginFilterFunc)
	beego.Router("/admin", &controllers.AdminIndexController{}, "*:Index")
	beego.Router("/admin/backenduser", &controllers.BackenduserController{}, "*:Index")

}
