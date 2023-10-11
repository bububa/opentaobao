package mtopopen

// PackageInfoVo 结构体
type PackageInfoVo struct {
	// 快递公司编码
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
	// 快递面单编码
	MailNo string `json:"mail_no,omitempty" xml:"mail_no,omitempty"`
	// 包裹id
	PackageId int64 `json:"package_id,omitempty" xml:"package_id,omitempty"`
}
