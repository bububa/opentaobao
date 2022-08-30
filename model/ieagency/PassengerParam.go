package ieagency

// PassengerParam 结构体
type PassengerParam struct {
	// 出生日期
	BirthDate string `json:"birth_date,omitempty" xml:"birth_date,omitempty"`
	// 证件持有人国籍名称
	DocHolderNationalityName string `json:"doc_holder_nationality_name,omitempty" xml:"doc_holder_nationality_name,omitempty"`
	// 证件号
	DocId string `json:"doc_id,omitempty" xml:"doc_id,omitempty"`
	// 证件发放国际名称
	DocIssueCountryName string `json:"doc_issue_country_name,omitempty" xml:"doc_issue_country_name,omitempty"`
	// 有效期截止日期
	EffectiveDate string `json:"effective_date,omitempty" xml:"effective_date,omitempty"`
	// 乘机人姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 证件类型（0：护照; 1:港澳通行证;2:台湾通行证;3:台胞证;4:回乡证;6:入台证）
	CertType int64 `json:"cert_type,omitempty" xml:"cert_type,omitempty"`
	// 乘机人性别（0:男;1:女）
	Gender int64 `json:"gender,omitempty" xml:"gender,omitempty"`
	// 乘机人类型(0:成人; 1 儿童)
	PassengerType int64 `json:"passenger_type,omitempty" xml:"passenger_type,omitempty"`
}
