package baichuan

// PasswordRuleResultDto 结构体
type PasswordRuleResultDto struct {
	// miaoPasswordRegular
	MiaoPasswordRegulars []string `json:"miao_password_regulars,omitempty" xml:"miao_password_regulars>string,omitempty"`
	// passwordRegular
	PasswordRegulars []string `json:"password_regulars,omitempty" xml:"password_regulars>string,omitempty"`
	// level
	Level string `json:"level,omitempty" xml:"level,omitempty"`
}
