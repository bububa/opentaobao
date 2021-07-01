package tmallhk

// CtsNgtc 结构体
type CtsNgtc struct {
	// ngtc证书编号
	ReportNo string `json:"report_no,omitempty" xml:"report_no,omitempty"`
	// ngtc证书防伪码
	ReportVerifyNo string `json:"report_verify_no,omitempty" xml:"report_verify_no,omitempty"`
}
