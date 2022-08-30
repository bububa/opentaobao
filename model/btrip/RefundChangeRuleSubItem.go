package btrip

// RefundChangeRuleSubItem 结构体
type RefundChangeRuleSubItem struct {
	// 退款子内容
	RefundSubContents []RefundChangeRuleSubContent `json:"refund_sub_contents,omitempty" xml:"refund_sub_contents>refund_change_rule_sub_content,omitempty"`
	// PTC
	Ptc string `json:"ptc,omitempty" xml:"ptc,omitempty"`
	// 退改签规则的类型
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 是否结构体
	IsStruct bool `json:"is_struct,omitempty" xml:"is_struct,omitempty"`
}
