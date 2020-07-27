package models

// ResponseResult 返回
type ResponseResult struct {
	Total int64       `json:"total"`
	Rows  interface{} `json:"rows"`
}

// JSONResult 返回结果
type JSONResult struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// type JsonResponse struct {
// }

//ResourceType 资源类型
type ResourceType = int64

const (
	PanelType   ResourceType = 1 //1：panel
	MenuType    ResourceType = 2 //2：菜单
	BtnResource ResourceType = 3 //3：按钮
)

const (
	STATUS_NORMAL   uint8 = 1 // 正常
	STATUS_ABNORMAL uint8 = 2
)
