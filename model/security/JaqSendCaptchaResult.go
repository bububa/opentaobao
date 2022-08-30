package security

// JaqSendCaptchaResult 结构体
type JaqSendCaptchaResult struct {
	// 扩展字段，格式为JSON字符串，由于出参“滚小球”等验证方式所需的额外出参，例如小球位置坐标等，请参考示例
	ExtendData string `json:"extend_data,omitempty" xml:"extend_data,omitempty"`
	// 验证码会话ID
	SessionId string `json:"session_id,omitempty" xml:"session_id,omitempty"`
	// 安全验证前向化下发参数结构体
	JaqDispatchParam *JaqDispatchParam `json:"jaq_dispatch_param,omitempty" xml:"jaq_dispatch_param,omitempty"`
	// 验证发起请求是否调用成功（及状态），约定正值为成功，负值为失败
	SendStatus int64 `json:"send_status,omitempty" xml:"send_status,omitempty"`
}
