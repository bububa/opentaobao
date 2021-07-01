package admarket

// AdSlot 结构体
type AdSlot struct {
	// 广告位id
	AdSlotId string `json:"ad_slot_id,omitempty" xml:"ad_slot_id,omitempty"`
	// 查询条件
	Query string `json:"query,omitempty" xml:"query,omitempty"`
	// 个数
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
}
