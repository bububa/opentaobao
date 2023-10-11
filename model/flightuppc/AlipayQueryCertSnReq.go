package flightuppc

// AlipayQueryCertSnReq 结构体
type AlipayQueryCertSnReq struct {
	// 业务代码
	ExternalAppletBizCode string `json:"external_applet_biz_code,omitempty" xml:"external_applet_biz_code,omitempty"`
	// 证书签名算法
	SignType string `json:"sign_type,omitempty" xml:"sign_type,omitempty"`
}
