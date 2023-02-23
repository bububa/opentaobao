package flight

// BaggageItemDto 结构体
type BaggageItemDto struct {
	// 行李业务对象集合
	BaggageBOList []BaggageDto `json:"baggage_b_o_list,omitempty" xml:"baggage_b_o_list>baggage_dto,omitempty"`
}
