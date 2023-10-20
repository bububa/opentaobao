package bus

import (
	"sync"
)

// BusNumberSearchRq 结构体
type BusNumberSearchRq struct {
	// 出发城市
	DepCityName string `json:"dep_city_name,omitempty" xml:"dep_city_name,omitempty"`
	// 出入日期
	DepDate string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	// 目的地
	LastPlaceName string `json:"last_place_name,omitempty" xml:"last_place_name,omitempty"`
}

var poolBusNumberSearchRq = sync.Pool{
	New: func() any {
		return new(BusNumberSearchRq)
	},
}

// GetBusNumberSearchRq() 从对象池中获取BusNumberSearchRq
func GetBusNumberSearchRq() *BusNumberSearchRq {
	return poolBusNumberSearchRq.Get().(*BusNumberSearchRq)
}

// ReleaseBusNumberSearchRq 释放BusNumberSearchRq
func ReleaseBusNumberSearchRq(v *BusNumberSearchRq) {
	v.DepCityName = ""
	v.DepDate = ""
	v.LastPlaceName = ""
	poolBusNumberSearchRq.Put(v)
}
