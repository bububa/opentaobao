package flightuppc

// ResultDo 结构体
type ResultDo struct {
	// 错误信息
	MsgForClient string `json:"msg_for_client,omitempty" xml:"msg_for_client,omitempty"`
	// 返回code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误信息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 是否回流成功
	Module bool `json:"module,omitempty" xml:"module,omitempty"`
	// 是否执行回流成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
