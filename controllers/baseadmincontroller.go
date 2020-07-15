package controllers

import (
	"strings"

	"github.com/xiaofan0208/xadmin/models"
)

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
	var tplName string
	switch len(tpl) {
	case 1:
		tplName = tpl[0]
		baselayout = baselayout
		break
	case 2:
		tplName = tpl[0]
		baselayout = tpl[1]
		break
	}
	if strings.TrimSpace(tplName) == "" {
		tplName = "admin/base/empty.html"
	}
	ctl.TplName = tplName
	ctl.Layout = baselayout
}
