package tmallservice

// TmallservicecenteranomalyrecoursehomedecorationquerybyidResult 结构体
type TmallservicecenteranomalyrecoursehomedecorationquerybyidResult struct {
	// 投诉单对象的json字符串
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
