package controllers

import (
	"github.com/astaxie/beego"
	"github.com/xiaofan0208/go-xadmin/models"
)

// MenuController 菜单管理
type MenuController struct {
	BaseAdminController
}

// Index index
func (ctl *MenuController) Index() {
	ctl.Data["PageName"] = "菜单"
	ctl.Data["PageDesc"] = "列表"
	// ctl.Data["ShowSearch"] = true // 是否显示搜索框
	// ctl.Data["canDelete"] = true  // 可删除
	// ctl.Data["canCreate"] = true  // 可新建

	ctl.Data["createURL"] = ctl.URLFor(".Create")
	ctl.SetTpl("", "admin/base/base_treelist_view.html")

	ctl.LayoutSections = make(map[string]string)
	ctl.LayoutSections["FooterScripts"] = "admin/rabc/menu/list_footerjs.html"
	//ctl.LayoutSections["LayoutSearch"] = "admin/rabc/menu/list_search.html"
}

// TreeGrid TreeGrid
func (ctl *MenuController) TreeGrid() {

	// 获取所有菜单
	param := make(map[string]interface{})
	param["Deleted"] = false
	param["Status"] = models.STATUS_NORMAL
	allmenus, total, err := models.GetMenuResourceByParam(param, 0, 0)
	if err != nil {
		beego.Error("models.GetMenuResourceByParam", err.Error())
	}
	tablelines := make([]interface{}, 0)
	for _, menu := range allmenus {
		one := make(map[string]interface{})
		one["id"] = menu.Id
		one["pid"] = menu.Pid
		one["title"] = menu.Title
		one["name"] = menu.Name
		one["type"] = menu.Type
		one["icon"] = menu.Icon
		one["link"] = ctl.URLFor(menu.UrlFor)
		tablelines = append(tablelines, one)
	}

	result := &models.ResponseResult{
		Total: total,
		Rows:  tablelines,
	}
	ctl.Data["json"] = result
	ctl.ServeJSON()

}
