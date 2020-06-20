package controllers

//BaseAdminController 后台控制器
type BaseAdminController struct {
	BaseController
}

// Prepare  prepare
func (ctl *BaseAdminController) Prepare() {
	ctl.BaseController.Prepare()

}
