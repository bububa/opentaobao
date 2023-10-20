package tmallsc

// TmallservicecenteranomalyrecoursehomedecorationcreateResult 结构体
type TmallservicecenteranomalyrecoursehomedecorationcreateResult struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 生成的投诉单id
	Value int64 `json:"value,omitempty" xml:"value,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
