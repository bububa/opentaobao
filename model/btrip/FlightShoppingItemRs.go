package btrip

import (
	"sync"
)

// FlightShoppingItemRs 结构体
type FlightShoppingItemRs struct {
	// 每一航段对应的仓位、价格
	SegmentCabinPrices []SegmentCabinPriceRs `json:"segment_cabin_prices,omitempty" xml:"segment_cabin_prices>segment_cabin_price_rs,omitempty"`
	// 搜索提供的报价
	SearchPrice *SearchPriceRs `json:"search_price,omitempty" xml:"search_price,omitempty"`
}

var poolFlightShoppingItemRs = sync.Pool{
	New: func() any {
		return new(FlightShoppingItemRs)
	},
}

// GetFlightShoppingItemRs() 从对象池中获取FlightShoppingItemRs
func GetFlightShoppingItemRs() *FlightShoppingItemRs {
	return poolFlightShoppingItemRs.Get().(*FlightShoppingItemRs)
}

// ReleaseFlightShoppingItemRs 释放FlightShoppingItemRs
func ReleaseFlightShoppingItemRs(v *FlightShoppingItemRs) {
	v.SegmentCabinPrices = v.SegmentCabinPrices[:0]
	v.SearchPrice = nil
	poolFlightShoppingItemRs.Put(v)
}
