package alitripmerchant

// VoucherVo 结构体
type VoucherVo struct {
	// 加价规则
	MarkUpRuleList []string `json:"mark_up_rule_list,omitempty" xml:"mark_up_rule_list>string,omitempty"`
	// 加价信息
	MarkUpInfo string `json:"mark_up_info,omitempty" xml:"mark_up_info,omitempty"`
	// 加价金额
	MarkUpAmount int64 `json:"mark_up_amount,omitempty" xml:"mark_up_amount,omitempty"`
	// 是否加价
	IsMarkUp bool `json:"is_mark_up,omitempty" xml:"is_mark_up,omitempty"`
}
