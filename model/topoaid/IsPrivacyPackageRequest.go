package topoaid

// IsPrivacyPackageRequest 结构体
type IsPrivacyPackageRequest struct {
	// 运单号
	MailNo string `json:"mail_no,omitempty" xml:"mail_no,omitempty"`
	// 快递公司编码
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
}
