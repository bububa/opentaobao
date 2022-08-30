package alitripmerchant

// UserAgreementVO 结构体
type UserAgreementVO struct {
	// 隐私协议
	ShowPrivacyAgreement bool `json:"show_privacy_agreement,omitempty" xml:"show_privacy_agreement,omitempty"`
	// 境外输出个人信息协议
	ShowDataExportAgreement bool `json:"show_data_export_agreement,omitempty" xml:"show_data_export_agreement,omitempty"`
}
