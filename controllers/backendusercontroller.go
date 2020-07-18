package controllers

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/xiaofan0208/xadmin/models"
)

// BackenduserController 管理员
type BackenduserController struct {
	BaseAdminController
}

// Index index
func (ctl *BackenduserController) Index() {
	ctl.Data["PageName"] = "管理员"
	ctl.Data["PageDesc"] = "列表"
	ctl.Data["ShowSearch"] = true // 是否显示搜索框

	ctl.SetTpl("", "admin/base/base_list_view.html")

	ctl.LayoutSections = make(map[string]string)
	ctl.LayoutSections["FooterScripts"] = "admin/user/list_footerjs.html"
	ctl.LayoutSections["LayoutSearch"] = "admin/user/list_search.html"
}

// PostList PostList
func (ctl *BackenduserController) PostList() {
	limit, _ := ctl.GetInt("limit", 0)
	offset, _ := ctl.GetInt("offset", 0)
	sort := ctl.GetString("sort")
	sortOrder := ctl.GetString("sortOrder")

	query := make(map[string]interface{})
	exclude := make(map[string]interface{})
	var orders []string

	// 关键字搜索
	search := ctl.GetString("search")
	params := make(map[string]interface{})
	if strings.TrimSpace(search) != "" {
		err := json.Unmarshal([]byte(search), &params)
		if err != nil {
			beego.Error("json.Unmarshal:", err.Error())
		}
	}

	if _, ok := params["id"]; ok {
		query["Id"] = params["id"]
	}
	if _, ok := params["name"]; ok {
		query["Name"] = params["name"]
	}

	query["Deleted"] = false
	// 排序
	if sortOrder == "" {
		sortOrder = "asc"
	}
	if sort != "" {
		if sortOrder == "desc" {
			sort = "-" + sort
		}
		orders = append(orders, sort)
	}

	result, err := ctl.querylist(query, exclude, orders, uint8(limit), uint8(offset))
	if err != nil {
		beego.Error("ctl.querylist:", err.Error())
	}

	ctl.Data["json"] = result
	ctl.ServeJSON()
}

func (ctl *BackenduserController) querylist(params map[string]interface{}, exclude map[string]interface{}, orders []string, limit uint8, offset uint8) (*models.ResponseResult, error) {
	records, total, err := models.GetAllBackenduserList(params, exclude, orders, limit, offset)
	if err != nil {
		beego.Error("models.GetAllBackenduserList", err.Error())
		return nil, err
	}
	tablelines := make([]interface{}, 0)
	for _, record := range records {
		one := make(map[string]interface{})
		one["Id"] = record.Id
		one["Name"] = record.Name
		one["Email"] = record.Email
		one["IsAdmin"] = record.IsAdmin
		tablelines = append(tablelines, one)
	}
	result := &models.ResponseResult{
		Total: total,
		Rows:  tablelines,
	}

	return result, err
}

// DeleteBatch 批量删除
func (ctl *BackenduserController) DeleteBatch() {
	ids := strings.TrimSpace(ctl.GetString("ids"))
	if ids == "" {
		ctl.ResponseError(nil)
		return
	}
	idsArr := strings.Split(ids, ",")
	for _, id := range idsArr {
		idInt64, _ := strconv.ParseInt(id, 10, 64)
		err := models.DeleteBackenduserByID(idInt64)
		if err != nil {
			beego.Error("models.DeleteBackenduserByID:", err.Error())
			ctl.ResponseError(err.Error())
			return
		}
	}
	ctl.ResponseSuccess(nil)
}

// Edit 编辑
func (ctl *BackenduserController) Edit() {
	if ctl.Ctx.Request.Method == "POST" {
		ctl.PostEdit()
		return
	}
	ctl.Data["PageName"] = "管理员"
	ctl.Data["PageDesc"] = "编辑"

	ctl.Data["canDelete"] = true
	ctl.Data["listURL"] = ctl.URLFor(".Index")
	ctl.Data["formURL"] = ctl.URLFor(".Edit")

	id := ctl.Ctx.Input.Param(":id")
	if idInt64, err := strconv.ParseInt(id, 10, 64); err != nil {
		ctl.Data["PageError"] = err.Error()
	} else {
		if record, err := models.GetBackenduserByID(idInt64); err != nil {
			ctl.Data["PageError"] = err.Error()
		} else {
			ctl.Data["Record"] = record
		}
	}

	ctl.SetTpl("admin/user/edit_form.html", "admin/base/base_edit_view.html")

	ctl.LayoutSections = make(map[string]string)
	ctl.LayoutSections["HeadCSS"] = "admin/user/edit_headcss.html"
	ctl.LayoutSections["FooterScripts"] = "admin/user/edit_footerjs.html"

}

//PostEdit PostEdit
func (ctl *BackenduserController) PostEdit() {
	record := &models.Backenduser{}
	if err := ctl.ParseForm(record); err != nil {
		beego.Error("ctl.ParseForm : ", err.Error())
		ctl.ResponseError(nil)
		return
	}

	if _, err := models.UpdateBackenduserByField(record, models.BackenduserUpdateFields...); err != nil {
		beego.Error("models.UpdateBackenduserByField: ", err.Error())
		ctl.ResponseError(nil)
		return
	}
	ctl.ResponseSuccess(nil)
}
