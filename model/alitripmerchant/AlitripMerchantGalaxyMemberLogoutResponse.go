package alitripmerchant

// AlitripmerchantgalaxymemberlogoutResponse 结构体
type AlitripmerchantgalaxymemberlogoutResponse struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 登出是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
