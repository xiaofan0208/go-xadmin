package routers

import (
	"github.com/xiaofan0208/xadmin/controllers"

	"github.com/astaxie/beego"
)

func init() {

}

// Router 后台路由
func Router() {
	beego.Router("/admin", &controllers.AdminIndexController{}, "*:Index")

	beego.Router("/admin/login", &controllers.SignInController{})

	beego.Router("/admin/backenduser", &controllers.BackenduserController{}, "*:Index")
}
