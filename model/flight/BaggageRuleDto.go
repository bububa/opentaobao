package flight

// BaggageRuleDto 结构体
type BaggageRuleDto struct {
	// 行李单元
	BaggageItem *BaggageItemDto `json:"baggage_item,omitempty" xml:"baggage_item,omitempty"`
}
