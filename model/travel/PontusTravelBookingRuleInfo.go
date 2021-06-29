package travel

// PontusTravelBookingRuleInfo 
type PontusTravelBookingRuleInfo struct {
    // fee_included：费用包含，跟团游必填； fee_excluded：费用不含，所有类目必填； order_info：预定须知； extra_cost：其他费用，预留；
    RuleType   string `json:"rule_type,omitempty" xml:"rule_type,omitempty"`
    // 1500个字
    RuleDesc   string `json:"rule_desc,omitempty" xml:"rule_desc,omitempty"`
    // 分点组织的规则说明
    RuleList   []string `json:"rule_list,omitempty" xml:"rule_list>string,omitempty"`
}
