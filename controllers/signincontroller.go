package controllers

import (
	"net/url"

	"github.com/astaxie/beego"
	"github.com/xiaofan0208/xadmin/models"
)

// SignInController 登录页面
type SignInController struct {
	BaseController
}

// Get 登录页面
func (ctl *SignInController) Get() {
	user := ctl.GetSession("User")
	if user != nil {
		ctl.Redirect("/admin", 302)
	}
	ctl.TplName = "admin/signin/login.html"
}

// Post 登录
func (ctl *SignInController) Post() {
	ctl.TplName = "admin/signin/login.html"

	name := ctl.GetString("email")
	password := ctl.GetString("password")

	if name == "" || password == "" {
		ctl.Redirect("/admin/login", 302)
	}

	var (
		user *models.Backenduser
		err  error
	)
	if user, err = models.CheckUserByName(name, password); err != nil {
		beego.Error("models.CheckUserByName", err.Error())
		ctl.Data["ErrorMsg"] = err.Error()
	} else {
		ctl.SetSession("User", user)
		// 登录成功后跳转
		redirectTo := ctl.GetString("redirect_to", "")
		redirectURL := "/admin"
		if redirectTo != "" {
			redirectURL, err = url.QueryUnescape(redirectTo)
			if err != nil {
				beego.Error(err.Error())
			}
		}
		//通过验证跳转到主界面
		ctl.Redirect(redirectURL, 302)
	}

}
