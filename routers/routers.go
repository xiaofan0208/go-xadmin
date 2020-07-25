package routers

import (
	"github.com/xiaofan0208/go-xadmin/controllers"
	"github.com/xiaofan0208/go-xadmin/models"

	"github.com/astaxie/beego"
)

func init() {
	InitMenu()
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

// InitMenu 初始化菜单
func InitMenu() {
	// 1.基本信息
	baseAdminInfo := &models.MenuResource{Title: "基本信息", Type: models.MenuType, Name: "baseAdminInfo", Icon: "far fa-circle"}
	// [1].管理员管理
	backenduser := &models.MenuResource{Title: "管理员管理", Type: models.MenuType, Name: "backenduser", Icon: "fas fa-tachometer-alt", UrlFor: "BackenduserController.Index"}
	backenduserQuery := &models.MenuResource{Title: "查询", Type: models.BtnResource, Name: "query", Icon: "far fa-circle"}
	backenduserAdd := &models.MenuResource{Title: "新增", Type: models.BtnResource, Name: "add", Icon: "far fa-circle"}
	backenduserEdit := &models.MenuResource{Title: "修改", Type: models.BtnResource, Name: "edit", Icon: "far fa-circle"}
	backenduserDel := &models.MenuResource{Title: "删除", Type: models.BtnResource, Name: "del", Icon: "far fa-circle"}
	backenduser.Children = []*models.MenuResource{backenduserQuery, backenduserAdd, backenduserEdit, backenduserDel}

	baseAdminInfo.Children = []*models.MenuResource{backenduser}

	models.AddMenus(baseAdminInfo)
}
