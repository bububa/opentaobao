package btrip

// FlightShoppingItemRs 结构体
type FlightShoppingItemRs struct {
	// 搜索提供的报价
	SearchPrice *SearchPriceRs `json:"search_price,omitempty" xml:"search_price,omitempty"`
	// 每一航段对应的仓位、价格
	SegmentCabinPrices []SegmentCabinPriceRs `json:"segment_cabin_prices,omitempty" xml:"segment_cabin_prices>segment_cabin_price_rs,omitempty"`
}
