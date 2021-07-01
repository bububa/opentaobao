package admarket

// AdSlots 结构体
type AdSlots struct {
	// 广告位id
	AdSlotId string `json:"ad_slot_id,omitempty" xml:"ad_slot_id,omitempty"`
	// 广告集合
	Ads []Ad `json:"ads,omitempty" xml:"ads>ad,omitempty"`
}
