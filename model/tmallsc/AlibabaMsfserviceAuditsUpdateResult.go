package tmallsc

// AlibabamsfserviceauditsupdateResult 结构体
type AlibabamsfserviceauditsupdateResult struct {
	// 失败原因
	SystemError string `json:"system_error,omitempty" xml:"system_error,omitempty"`
	// 失败原因
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
