package baichuan

// PasswordRuleResultDto 
type PasswordRuleResultDto struct {

    // level
    Level   string `json:"level,omitempty"`

    // miaoPasswordRegular
    MiaoPasswordRegulars   []String `json:"miao_password_regulars,omitempty"`

    // passwordRegular
    PasswordRegulars   []String `json:"password_regulars,omitempty"`

}
