package axintrade

import (
	"sync"
)

// HotelArrivalTime 结构体
type HotelArrivalTime struct {
	// 最晚到店时间
	LatestArrivalTime string `json:"latest_arrival_time,omitempty" xml:"latest_arrival_time,omitempty"`
	// 最早到店时间
	EarliestArrivalTime string `json:"earliest_arrival_time,omitempty" xml:"earliest_arrival_time,omitempty"`
}

var poolHotelArrivalTime = sync.Pool{
	New: func() any {
		return new(HotelArrivalTime)
	},
}

// GetHotelArrivalTime() 从对象池中获取HotelArrivalTime
func GetHotelArrivalTime() *HotelArrivalTime {
	return poolHotelArrivalTime.Get().(*HotelArrivalTime)
}

// ReleaseHotelArrivalTime 释放HotelArrivalTime
func ReleaseHotelArrivalTime(v *HotelArrivalTime) {
	v.LatestArrivalTime = ""
	v.EarliestArrivalTime = ""
	poolHotelArrivalTime.Put(v)
}
