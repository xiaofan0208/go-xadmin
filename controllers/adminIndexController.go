package controllers

// AdminIndexController 后台首页
type AdminIndexController struct {
	BaseAdminController
}

// Index index
func (ctl *AdminIndexController) Index() {

	ctl.Layout = "admin/base/layout.html"
	ctl.TplName = "admin/index.html"

}
