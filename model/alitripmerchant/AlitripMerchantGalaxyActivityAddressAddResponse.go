package alitripmerchant

// AlitripmerchantgalaxyactivityaddressaddResponse 结构体
type AlitripmerchantgalaxyactivityaddressaddResponse struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误代码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 邮寄信息是否填写成功
	Content bool `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
