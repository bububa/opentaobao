package axindata

import (
	"sync"
)

// HourRoomInfo 结构体
type HourRoomInfo struct {
	// 小时房可入住开始时间，HH:mm格式
	HourRoomCheckInStart string `json:"hour_room_check_in_start,omitempty" xml:"hour_room_check_in_start,omitempty"`
	// 小时房可入住结束时间，HH:mm格式
	HourRoomCheckInEnd string `json:"hour_room_check_in_end,omitempty" xml:"hour_room_check_in_end,omitempty"`
	// 小时房入住时长，单位小时
	Hourage int64 `json:"hourage,omitempty" xml:"hourage,omitempty"`
}

var poolHourRoomInfo = sync.Pool{
	New: func() any {
		return new(HourRoomInfo)
	},
}

// GetHourRoomInfo() 从对象池中获取HourRoomInfo
func GetHourRoomInfo() *HourRoomInfo {
	return poolHourRoomInfo.Get().(*HourRoomInfo)
}

// ReleaseHourRoomInfo 释放HourRoomInfo
func ReleaseHourRoomInfo(v *HourRoomInfo) {
	v.HourRoomCheckInStart = ""
	v.HourRoomCheckInEnd = ""
	v.Hourage = 0
	poolHourRoomInfo.Put(v)
}
