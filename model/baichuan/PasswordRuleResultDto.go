package baichuan

// PasswordRuleResultDto 
type PasswordRuleResultDto struct {
    // level
    Level   string `json:"level,omitempty" xml:"level,omitempty"`
    // miaoPasswordRegular
    MiaoPasswordRegulars   []string `json:"miao_password_regulars,omitempty" xml:"miao_password_regulars>string,omitempty"`
    // passwordRegular
    PasswordRegulars   []string `json:"password_regulars,omitempty" xml:"password_regulars>string,omitempty"`
}
