package controllers

// MenuController 菜单管理
type MenuController struct {
	BaseAdminController
}

// Index index
func (ctl *MenuController) Index() {
	ctl.Data["PageName"] = "菜单"
	ctl.Data["PageDesc"] = "列表"
	ctl.Data["ShowSearch"] = true // 是否显示搜索框
	ctl.Data["canDelete"] = true  // 可删除
	ctl.Data["canCreate"] = true  // 可新建

	ctl.Data["createURL"] = ctl.URLFor(".Create")
	ctl.SetTpl("", "admin/base/base_list_view.html")

	ctl.LayoutSections = make(map[string]string)
	ctl.LayoutSections["FooterScripts"] = "admin/rabc/menu/list_footerjs.html"
	//ctl.LayoutSections["LayoutSearch"] = "admin/rabc/menu/list_search.html"
}
