package topoaid

// IsPrivacyPackageRequest 结构体
type IsPrivacyPackageRequest struct {
	// 运单号
	MailNo string `json:"mail_no,omitempty" xml:"mail_no,omitempty"`
	// 快递公司编码
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
	// 柜机黑名单手机号前7位，如有多个用英文逗号,分隔，每个元素仅前七位有效
	BlackMobiles string `json:"black_mobiles,omitempty" xml:"black_mobiles,omitempty"`
}
