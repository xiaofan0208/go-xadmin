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
