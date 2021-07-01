package admarket

// BidResponse 结构体
type BidResponse struct {
	// 广告位列表
	AdSlots []AdSlots `json:"ad_slots,omitempty" xml:"ad_slots>ad_slots,omitempty"`
}
