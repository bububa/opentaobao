package bus

import (
	"sync"
)

// B2BBusLineInfo 结构体
type B2BBusLineInfo struct {
	// 出发时间
	DepTime string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// 目的地
	LastPlaceName string `json:"last_place_name,omitempty" xml:"last_place_name,omitempty"`
	// 车次ID
	ScheduleId string `json:"schedule_id,omitempty" xml:"schedule_id,omitempty"`
	// 出发城市
	StartCityName string `json:"start_city_name,omitempty" xml:"start_city_name,omitempty"`
	// 出发车站
	StartStationName string `json:"start_station_name,omitempty" xml:"start_station_name,omitempty"`
	// 出发车站地址
	StartStationNameAddress string `json:"start_station_name_address,omitempty" xml:"start_station_name_address,omitempty"`
	// 到达车站
	ToStationName string `json:"to_station_name,omitempty" xml:"to_station_name,omitempty"`
}

var poolB2BBusLineInfo = sync.Pool{
	New: func() any {
		return new(B2BBusLineInfo)
	},
}

// GetB2BBusLineInfo() 从对象池中获取B2BBusLineInfo
func GetB2BBusLineInfo() *B2BBusLineInfo {
	return poolB2BBusLineInfo.Get().(*B2BBusLineInfo)
}

// ReleaseB2BBusLineInfo 释放B2BBusLineInfo
func ReleaseB2BBusLineInfo(v *B2BBusLineInfo) {
	v.DepTime = ""
	v.LastPlaceName = ""
	v.ScheduleId = ""
	v.StartCityName = ""
	v.StartStationName = ""
	v.StartStationNameAddress = ""
	v.ToStationName = ""
	poolB2BBusLineInfo.Put(v)
}
