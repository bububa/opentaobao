package wdk

// ShopRangeUpdateRequest 结构体
type ShopRangeUpdateRequest struct {
	// 销售范围
	ShopRanges []ShopRange `json:"shop_ranges,omitempty" xml:"shop_ranges>shop_range,omitempty"`
	// 经营店ID
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 渠道
	ChannelSourceType int64 `json:"channel_source_type,omitempty" xml:"channel_source_type,omitempty"`
}
