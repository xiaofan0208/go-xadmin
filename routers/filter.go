package routers

import (
	"net/url"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/xiaofan0208/go-xadmin/models"
)

// LoginFilterFunc 登陆前session判断
var LoginFilterFunc = func(ctx *context.Context) {
	if strings.Contains(ctx.Request.RequestURI, "/admin/login") {
		return
	}

	user := ctx.Input.Session("User")
	if user == nil {
		redirectURL := "/admin/login"
		redirectTo := ctx.Request.RequestURI
		if redirectTo != "" {
			redirectTo = url.QueryEscape(redirectTo)
			redirectURL += "?redirect_to=" + redirectTo
		}
		ctx.Redirect(302, redirectURL)
	}
}

// MenuFilterFunc 菜单
var MenuFilterFunc = func(ctx *context.Context) {
	redirectTo := ctx.Request.RequestURI
	// 获取所有菜单
	param := make(map[string]interface{})
	param["Deleted"] = false
	param["Status"] = models.STATUS_NORMAL
	param["Type"] = models.MenuType
	allmenus, _, err := models.GetMenuResourceByParam(param, 0, 0)
	if err != nil {
		return
	}

	base := &beego.Controller{}
	for _, m := range allmenus {
		if m.UrlFor != "" {
			m.Link = base.URLFor(m.UrlFor)
		}
		if m.Link == "" {
			continue
		}
		if redirectTo == m.Link || strings.Contains(redirectTo, m.Link) {
			m.Active = true
		}
	}

	result := make([]*models.MenuResource, 0)
	for _, menu := range allmenus {
		if menu.Pid == 0 {
			children := resourceAddSons(menu, allmenus, result)
			menu.Children = children
			result = append(result, menu)
		}
	}

	ctx.Input.SetData("Menus", result)
}

func resourceAddSons(cur *models.MenuResource, list, result []*models.MenuResource) []*models.MenuResource {
	for _, item := range list {
		if item.Pid == cur.Id {
			children := resourceAddSons(item, list, item.Children)
			item.Children = children
			if item.Active == true {
				cur.Active = true
			}
			result = append(result, item)
		}
	}
	return result
}
