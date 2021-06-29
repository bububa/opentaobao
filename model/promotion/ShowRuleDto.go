package promotion

// ShowRuleDTO 
type ShowRuleDTO struct {
    // 规则扩展信息
    Feature   string `json:"feature,omitempty" xml:"feature,omitempty"`
    // 规则是否通过
    Passed   bool `json:"passed,omitempty" xml:"passed,omitempty"`
    // 规则类型
    Type   string `json:"type,omitempty" xml:"type,omitempty"`
}
