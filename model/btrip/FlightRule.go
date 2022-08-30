package btrip

// FlightRule 结构体
type FlightRule struct {
	// 行李额描述
	BaggageInfo string `json:"baggage_info,omitempty" xml:"baggage_info,omitempty"`
	// 额外信息
	Extra string `json:"extra,omitempty" xml:"extra,omitempty"`
	// 退改签信息
	TuigaiqianInfo string `json:"tuigaiqian_info,omitempty" xml:"tuigaiqian_info,omitempty"`
	// 改签规则
	ChangeRule *TgqNodeDo `json:"change_rule,omitempty" xml:"change_rule,omitempty"`
	// 退票规则
	RefundRule *TgqNodeDo `json:"refund_rule,omitempty" xml:"refund_rule,omitempty"`
	// 签转规则
	SignRule *TgqNodeDo `json:"sign_rule,omitempty" xml:"sign_rule,omitempty"`
	// 升舱规则
	UpgradeRule *TgqNodeDo `json:"upgrade_rule,omitempty" xml:"upgrade_rule,omitempty"`
	// 改签规则
	ChangeRuleItem *RefundChangeRuleItem `json:"change_rule_item,omitempty" xml:"change_rule_item,omitempty"`
	// 退票规则
	RefundRuleItem *RefundChangeRuleItem `json:"refund_rule_item,omitempty" xml:"refund_rule_item,omitempty"`
	// 行李规则
	BaggageItem *BaggageItem `json:"baggage_item,omitempty" xml:"baggage_item,omitempty"`
}
