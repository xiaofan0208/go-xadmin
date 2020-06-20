package controllers

import "github.com/astaxie/beego"

// BaseController 基础控制器
type BaseController struct {
	beego.Controller
}

// Prepare  prepare
func (ctl *BaseController) Prepare() {

}
