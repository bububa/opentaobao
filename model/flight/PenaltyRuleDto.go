package flight

// PenaltyRuleDto 结构体
type PenaltyRuleDto struct {
	// 退改签规则单元
	PenaltyItem *PenaltyItemDto `json:"penalty_item,omitempty" xml:"penalty_item,omitempty"`
}
