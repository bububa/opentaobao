package btrip

import (
	"sync"
)

// FlightSearchRs 结构体
type FlightSearchRs struct {
	// 组合商品列表
	ItemList []GroupItemRs `json:"item_list,omitempty" xml:"item_list>group_item_rs,omitempty"`
}

var poolFlightSearchRs = sync.Pool{
	New: func() any {
		return new(FlightSearchRs)
	},
}

// GetFlightSearchRs() 从对象池中获取FlightSearchRs
func GetFlightSearchRs() *FlightSearchRs {
	return poolFlightSearchRs.Get().(*FlightSearchRs)
}

// ReleaseFlightSearchRs 释放FlightSearchRs
func ReleaseFlightSearchRs(v *FlightSearchRs) {
	v.ItemList = v.ItemList[:0]
	poolFlightSearchRs.Put(v)
}
