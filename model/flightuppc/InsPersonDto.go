package flightuppc

// InsPersonDto 结构体
type InsPersonDto struct {
	// 证件编号
	CertNo string `json:"cert_no,omitempty" xml:"cert_no,omitempty"`
	// 证件名字
	CertName string `json:"cert_name,omitempty" xml:"cert_name,omitempty"`
	// 证件类型
	CertType int64 `json:"cert_type,omitempty" xml:"cert_type,omitempty"`
}
