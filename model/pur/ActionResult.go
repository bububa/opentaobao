package pur

// ActionResult 结构体
type ActionResult struct {
	// 系统自动生成
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 系统自动生成
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 系统自动生成
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// https://
	RedirectUrl string `json:"redirect_url,omitempty" xml:"redirect_url,omitempty"`
	// 返回结果
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 返回提示信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码
	Error string `json:"error,omitempty" xml:"error,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 返回值
	RetValue string `json:"ret_value,omitempty" xml:"ret_value,omitempty"`
}
