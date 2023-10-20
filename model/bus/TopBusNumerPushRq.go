package bus

import (
	"sync"
)

// TopBusNumerPushRq 结构体
type TopBusNumerPushRq struct {
	// 目的地
	LastPlaceName string `json:"last_place_name,omitempty" xml:"last_place_name,omitempty"`
	// 出发日期
	DepDate string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	// 出发城市
	FromCityName string `json:"from_city_name,omitempty" xml:"from_city_name,omitempty"`
}

var poolTopBusNumerPushRq = sync.Pool{
	New: func() any {
		return new(TopBusNumerPushRq)
	},
}

// GetTopBusNumerPushRq() 从对象池中获取TopBusNumerPushRq
func GetTopBusNumerPushRq() *TopBusNumerPushRq {
	return poolTopBusNumerPushRq.Get().(*TopBusNumerPushRq)
}

// ReleaseTopBusNumerPushRq 释放TopBusNumerPushRq
func ReleaseTopBusNumerPushRq(v *TopBusNumerPushRq) {
	v.LastPlaceName = ""
	v.DepDate = ""
	v.FromCityName = ""
	poolTopBusNumerPushRq.Put(v)
}
