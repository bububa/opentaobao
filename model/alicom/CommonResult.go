package alicom

// CommonResult 结构体
type CommonResult struct {
	// model
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// desc
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
