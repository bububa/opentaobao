package axintrade

// TravelerDto 结构体
type TravelerDto struct {
	// 姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 联系人手机号
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 联系人邮箱
	Email string `json:"email,omitempty" xml:"email,omitempty"`
	// 联系人证件号
	Certificates string `json:"certificates,omitempty" xml:"certificates,omitempty"`
	// 联系人证件类型
	CertificatesType int64 `json:"certificates_type,omitempty" xml:"certificates_type,omitempty"`
}
