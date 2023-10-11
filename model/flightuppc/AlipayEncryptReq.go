package flightuppc

// AlipayEncryptReq 结构体
type AlipayEncryptReq struct {
	// 待加密明文
	SourceContent string `json:"source_content,omitempty" xml:"source_content,omitempty"`
	// 业务代码
	ExternalAppletBizCode string `json:"external_applet_biz_code,omitempty" xml:"external_applet_biz_code,omitempty"`
}
