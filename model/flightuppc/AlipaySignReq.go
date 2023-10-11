package flightuppc

// AlipaySignReq 结构体
type AlipaySignReq struct {
	// 加签自定义参数，格式遵循http请求路径参数的格式
	SourceContent string `json:"source_content,omitempty" xml:"source_content,omitempty"`
	// 业务代码
	ExternalAppletBizCode string `json:"external_applet_biz_code,omitempty" xml:"external_applet_biz_code,omitempty"`
	// 签名算法类型
	SignType string `json:"sign_type,omitempty" xml:"sign_type,omitempty"`
}
