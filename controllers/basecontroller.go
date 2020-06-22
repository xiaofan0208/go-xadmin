package controllers

import (
	"github.com/astaxie/beego"
	"github.com/xiaofan0208/xadmin/models"
)

// BaseController 基础控制器
type BaseController struct {
	beego.Controller
	User *models.Backenduser
}

// Prepare  prepare
func (ctl *BaseController) Prepare() {

}
