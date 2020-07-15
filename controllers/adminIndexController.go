package controllers

// AdminIndexController 后台首页
type AdminIndexController struct {
	BaseAdminController
}

// Index index
func (ctl *AdminIndexController) Index() {

	ctl.SetTpl("admin/index.html", "admin/base/layout.html")

	ctl.LayoutSections = make(map[string]string)
	ctl.LayoutSections["HeadCSS"] = "admin/index_headcss.html"
	ctl.LayoutSections["FooterScripts"] = "admin/index_footerjs.html"

}
