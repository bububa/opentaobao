package alihouse

// ExistingHomeElectricAgreementDto 结构体
type ExistingHomeElectricAgreementDto struct {
	// 协议内容
	AgreementContent string `json:"agreement_content,omitempty" xml:"agreement_content,omitempty"`
	// 协议名称
	AgreementName string `json:"agreement_name,omitempty" xml:"agreement_name,omitempty"`
	// 外部协议ID
	OuterAgreementId string `json:"outer_agreement_id,omitempty" xml:"outer_agreement_id,omitempty"`
	// 是否生效 1-生效 0-失效
	IsValid int64 `json:"is_valid,omitempty" xml:"is_valid,omitempty"`
	// 是否使用模板 1-是 0-否
	IsTemplate int64 `json:"is_template,omitempty" xml:"is_template,omitempty"`
	// 协议类型
	AgreementType int64 `json:"agreement_type,omitempty" xml:"agreement_type,omitempty"`
	// 公司进件ID
	MerchantOpenId int64 `json:"merchant_open_id,omitempty" xml:"merchant_open_id,omitempty"`
}
