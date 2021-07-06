package util

// CommonResult 结构体
type CommonResult struct {
	// 服务子错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 服务子扩展错误码
	SubCode string `json:"sub_code,omitempty" xml:"sub_code,omitempty"`
	// 描述信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 描述错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 描述请求信息
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 无
	Model *AlibabaTaobaoWtUserCrowdModel `json:"model,omitempty" xml:"model,omitempty"`
	// 成功标识
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
