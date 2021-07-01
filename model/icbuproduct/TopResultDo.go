package icbuproduct

// TopResultDo 结构体
type TopResultDo struct {
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 库存更新是否成功
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 接口更新失败时错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 失败错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 接口失败时的追踪id
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
}
