package alicom

// CommonResult 结构体
type CommonResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// desc
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 阿里订单号
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 接口返回成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
