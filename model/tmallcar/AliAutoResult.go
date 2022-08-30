package tmallcar

// AliAutoResult 结构体
type AliAutoResult struct {
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 错误code
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 结果
	Data *DaSouEticketVerifyResultDto `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
