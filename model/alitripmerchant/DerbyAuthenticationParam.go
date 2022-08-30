package alitripmerchant

// DerbyAuthenticationParam 结构体
type DerbyAuthenticationParam struct {
	// 认证邮箱
	Email string `json:"email,omitempty" xml:"email,omitempty"`
	// 认证密码
	Password string `json:"password,omitempty" xml:"password,omitempty"`
	// 用户注册类型
	UserRegisterType string `json:"user_register_type,omitempty" xml:"user_register_type,omitempty"`
	// 是否接受会员协议
	AcceptedTandC string `json:"accepted_tand_c,omitempty" xml:"accepted_tand_c,omitempty"`
	// 是否是否勾选ALL和SMSALL收取优惠邮件
	OptinAll bool `json:"optin_all,omitempty" xml:"optin_all,omitempty"`
	// 是否同意向境外提供个人信息
	DataExportAgreement bool `json:"data_export_agreement,omitempty" xml:"data_export_agreement,omitempty"`
}
