package ioti

// AgentConfigResponse 结构体
type AgentConfigResponse struct {
	// 调用方法
	Method string `json:"method,omitempty" xml:"method,omitempty"`
	// 返回编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 返回消息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 消息体
	Result *ConfigResult `json:"result,omitempty" xml:"result,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
