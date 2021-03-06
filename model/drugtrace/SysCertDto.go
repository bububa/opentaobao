package drugtrace

// SysCertDto 结构体
type SysCertDto struct {
	// 证书序列号
	CertSn string `json:"cert_sn,omitempty" xml:"cert_sn,omitempty"`
	// 证书公钥
	Cert string `json:"cert,omitempty" xml:"cert,omitempty"`
}
