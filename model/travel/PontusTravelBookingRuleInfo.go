package travel

// PontusTravelBookingRuleInfo 结构体
type PontusTravelBookingRuleInfo struct {
	// 分点组织的规则说明
	RuleList []string `json:"rule_list,omitempty" xml:"rule_list>string,omitempty"`
	// Fee_Included：费用包含；  Order_Info：预定须知；
	RuleType string `json:"rule_type,omitempty" xml:"rule_type,omitempty"`
	// 1500个字符
	RuleDesc string `json:"rule_desc,omitempty" xml:"rule_desc,omitempty"`
}
