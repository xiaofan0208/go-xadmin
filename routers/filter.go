package routers

import (
	"net/url"
	"strings"

	"github.com/astaxie/beego/context"
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
