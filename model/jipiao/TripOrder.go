package jipiao

import (
	"sync"
)

// TripOrder 结构体
type TripOrder struct {
	// 订单航班信息（包含航班乘机人）
	FlightInfos []TripFlightInfo `json:"flight_infos,omitempty" xml:"flight_infos>trip_flight_info,omitempty"`
	// 扩展字段
	Extra string `json:"extra,omitempty" xml:"extra,omitempty"`
	// 订单基础信息
	BaseInfo *TripBaseInfo `json:"base_info,omitempty" xml:"base_info,omitempty"`
	// 订单行程单信息
	Itinerary *Itinerary `json:"itinerary,omitempty" xml:"itinerary,omitempty"`
	// 订单行政购票信息
	CorpInfo *CorpInfo `json:"corp_info,omitempty" xml:"corp_info,omitempty"`
}

var poolTripOrder = sync.Pool{
	New: func() any {
		return new(TripOrder)
	},
}

// GetTripOrder() 从对象池中获取TripOrder
func GetTripOrder() *TripOrder {
	return poolTripOrder.Get().(*TripOrder)
}

// ReleaseTripOrder 释放TripOrder
func ReleaseTripOrder(v *TripOrder) {
	v.FlightInfos = v.FlightInfos[:0]
	v.Extra = ""
	v.BaseInfo = nil
	v.Itinerary = nil
	v.CorpInfo = nil
	poolTripOrder.Put(v)
}
