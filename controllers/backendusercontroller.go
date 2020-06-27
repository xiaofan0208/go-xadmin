package controllers

// BackenduserController 管理员
type BackenduserController struct {
	BaseAdminController
}

// Index index
func (ctl *BackenduserController) Index() {
	ctl.Data["PageName"] = "管理员列表"
	ctl.SetTpl("admin/user/list.html", "admin/base/layout.html")

	ctl.LayoutSections = make(map[string]string)
	ctl.LayoutSections["HtmlHead"] = "admin/user/list_headcss.html"
	ctl.LayoutSections["FooterScripts"] = "admin/user/list_footerjs.html"

}
