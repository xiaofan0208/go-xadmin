package controllers

import (
	"github.com/astaxie/beego"
	"github.com/xiaofan0208/go-xadmin/models"
	xbaseModels "github.com/xiaofan0208/go-xbase/models"
)

// BaseController 基础控制器
type BaseController struct {
	beego.Controller
	User *models.Backenduser
}

// Prepare  prepare
func (ctl *BaseController) Prepare() {

}

// ResponseList 返回结果
func (ctl *BaseController) ResponseList(code int, msg string, data interface{}) {
	result := &models.JSONResult{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	ctl.Data["json"] = result
	ctl.ServeJSON()
}

// ResponseSuccess 返回成功
func (ctl *BaseController) ResponseSuccess(data interface{}) {
	result := &models.JSONResult{
		Code: int(xbaseModels.RECODE_OK),
		Msg:  xbaseModels.RecodeText(xbaseModels.RECODE_OK),
		Data: data,
	}
	ctl.Data["json"] = result
	ctl.ServeJSON()
	ctl.StopRun()
}

// ResponseError 返回失败
func (ctl *BaseController) ResponseError(msg string) {
	result := &models.JSONResult{
		Code: -1,
		Msg:  msg,
		Data: nil,
	}
	ctl.Data["json"] = result
	ctl.ServeJSON()
	ctl.StopRun()
}
