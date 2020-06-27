package routers

import (
	"github.com/astaxie/beego/context"
)

// LoginFilterFunc 登陆前session判断
var LoginFilterFunc = func(ctx *context.Context) {
	if ctx.Request.RequestURI == "/admin/login" {
		return
	}

	user := ctx.Input.Session("User")
	if user == nil {
		ctx.Redirect(302, "/admin/login")
	}
}
