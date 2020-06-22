package controllers

// AdminIndexController 后台首页
type AdminIndexController struct {
	BaseAdminController
}

// Index index
func (ctl *AdminIndexController) Index() {

	ctl.SetTpl("admin/index.html", "admin/base/layout.html")

}
