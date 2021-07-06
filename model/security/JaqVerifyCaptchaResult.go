package security

// JaqVerifyCaptchaResult 结构体
type JaqVerifyCaptchaResult struct {
	// 扩展字段，格式为JSON字符串，当前为预留字段
	ExtendData string `json:"extend_data,omitempty" xml:"extend_data,omitempty"`
	// 安全验证前向化下发参数结构体
	JaqDispatchParam *JaqDispatchParam `json:"jaq_dispatch_param,omitempty" xml:"jaq_dispatch_param,omitempty"`
	// 验证检查请求是否调用成功（及状态），约定正值为成功，负值为失败
	VerifyStatus int64 `json:"verify_status,omitempty" xml:"verify_status,omitempty"`
}
