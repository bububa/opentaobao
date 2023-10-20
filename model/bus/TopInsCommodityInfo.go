package bus

import (
	"sync"
)

// TopInsCommodityInfo 结构体
type TopInsCommodityInfo struct {
	// 出行日期，精确到分钟 yyyyMMddHHmm
	TravelDate string `json:"travel_date,omitempty" xml:"travel_date,omitempty"`
	// 起始站点
	StartStationId string `json:"start_station_id,omitempty" xml:"start_station_id,omitempty"`
	// 票价，单位：分
	TicketPrice int64 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
	// 行程时长，精确到分钟
	ItineraryTime int64 `json:"itinerary_time,omitempty" xml:"itinerary_time,omitempty"`
}

var poolTopInsCommodityInfo = sync.Pool{
	New: func() any {
		return new(TopInsCommodityInfo)
	},
}

// GetTopInsCommodityInfo() 从对象池中获取TopInsCommodityInfo
func GetTopInsCommodityInfo() *TopInsCommodityInfo {
	return poolTopInsCommodityInfo.Get().(*TopInsCommodityInfo)
}

// ReleaseTopInsCommodityInfo 释放TopInsCommodityInfo
func ReleaseTopInsCommodityInfo(v *TopInsCommodityInfo) {
	v.TravelDate = ""
	v.StartStationId = ""
	v.TicketPrice = 0
	v.ItineraryTime = 0
	poolTopInsCommodityInfo.Put(v)
}
