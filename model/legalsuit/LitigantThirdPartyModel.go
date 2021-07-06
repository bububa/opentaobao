package legalsuit

// LitigantThirdPartyModel 结构体
type LitigantThirdPartyModel struct {
	// 承办律师联系方式
	LawyerContact string `json:"lawyer_contact,omitempty" xml:"lawyer_contact,omitempty"`
	// 承办律师姓名
	LawyerName string `json:"lawyer_name,omitempty" xml:"lawyer_name,omitempty"`
	// 承办律所名称
	LawFirmName string `json:"law_firm_name,omitempty" xml:"law_firm_name,omitempty"`
	// 是否为集团公司
	IsAlibabaCompany string `json:"is_alibaba_company,omitempty" xml:"is_alibaba_company,omitempty"`
	// 证件编号
	CertifyNumber string `json:"certify_number,omitempty" xml:"certify_number,omitempty"`
	// 证件类型
	CertifyType string `json:"certify_type,omitempty" xml:"certify_type,omitempty"`
	// 联系方式
	Contact string `json:"contact,omitempty" xml:"contact,omitempty"`
	// 名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 序号
	SerialNumber int64 `json:"serial_number,omitempty" xml:"serial_number,omitempty"`
	// 住所地
	Address bool `json:"address,omitempty" xml:"address,omitempty"`
}
