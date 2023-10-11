package flightuppc

// AlipayCertSnDto 结构体
type AlipayCertSnDto struct {
	// app证书序列号
	AppCertSn string `json:"app_cert_sn,omitempty" xml:"app_cert_sn,omitempty"`
	// alipay证书序列号
	AlipayCertSn string `json:"alipay_cert_sn,omitempty" xml:"alipay_cert_sn,omitempty"`
	// alipay根证书序列号
	AlipayRootCertSn string `json:"alipay_root_cert_sn,omitempty" xml:"alipay_root_cert_sn,omitempty"`
}
