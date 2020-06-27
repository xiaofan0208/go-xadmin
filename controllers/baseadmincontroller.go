package controllers

import "github.com/xiaofan0208/xadmin/models"

//BaseAdminController 后台控制器
type BaseAdminController struct {
	BaseController
}

// Prepare  prepare
func (ctl *BaseAdminController) Prepare() {
	ctl.BaseController.Prepare()

	user := ctl.GetSession("User")
	if user != nil {
		ctl.User = user.(*models.Backenduser)
		ctl.Data["User"] = ctl.User
	}
}

// SetTpl 设置布局文件
func (ctl *BaseAdminController) SetTpl(tpl ...string) {
	baselayout := "admin/base/layout.html"
	switch len(tpl) {
	case 1:
		ctl.TplName = tpl[0]
		ctl.Layout = baselayout
		break
	case 2:
		ctl.TplName = tpl[0]
		ctl.Layout = tpl[1]
		break
	}
}
