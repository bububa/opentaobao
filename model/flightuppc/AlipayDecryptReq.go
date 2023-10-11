package flightuppc

// AlipayDecryptReq 结构体
type AlipayDecryptReq struct {
	// 待解密密文
	EncryptContent string `json:"encrypt_content,omitempty" xml:"encrypt_content,omitempty"`
	// 业务代码
	ExternalAppletBizCode string `json:"external_applet_biz_code,omitempty" xml:"external_applet_biz_code,omitempty"`
}
