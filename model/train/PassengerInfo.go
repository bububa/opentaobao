package train

// PassengerInfo 结构体
type PassengerInfo struct {
	// 证件类型
	CertificateTypeCode string `json:"certificate_type_code,omitempty" xml:"certificate_type_code,omitempty"`
	// 乘客
	PassengerName string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// 证件号
	CertificateNo string `json:"certificate_no,omitempty" xml:"certificate_no,omitempty"`
}
