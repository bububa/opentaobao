package promotion

// SingleRuleDTO 
type SingleRuleDTO struct {
    // 规则CODE
    RuleCode   string `json:"rule_code,omitempty" xml:"rule_code,omitempty"`
    // 规则参数
    Params   string `json:"params,omitempty" xml:"params,omitempty"`
}
