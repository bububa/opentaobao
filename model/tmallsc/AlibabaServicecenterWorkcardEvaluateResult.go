package tmallsc

// AlibabaservicecenterworkcardevaluateResult 结构体
type AlibabaservicecenterworkcardevaluateResult struct {
	// 错误原因文案
	DisplayMsg string `json:"display_msg,omitempty" xml:"display_msg,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误原因
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 是否调用成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
