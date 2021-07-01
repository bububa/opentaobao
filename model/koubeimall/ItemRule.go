package koubeimall

// ItemRule 结构体
type ItemRule struct {
	// 规则列表
	RuleList []string `json:"rule_list,omitempty" xml:"rule_list>string,omitempty"`
}
