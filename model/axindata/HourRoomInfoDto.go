package axindata

import (
	"sync"
)

// HourRoomInfoDto 结构体
type HourRoomInfoDto struct {
	// 小时房可入住开始时间，HH:mm格式
	HourRoomCheckInStart string `json:"hour_room_check_in_start,omitempty" xml:"hour_room_check_in_start,omitempty"`
	// 小时房可入住结束时间，HH:mm格式
	HourRoomCheckInEnd string `json:"hour_room_check_in_end,omitempty" xml:"hour_room_check_in_end,omitempty"`
	// 小时房入住时长，单位小时
	Hourage int64 `json:"hourage,omitempty" xml:"hourage,omitempty"`
}

var poolHourRoomInfoDto = sync.Pool{
	New: func() any {
		return new(HourRoomInfoDto)
	},
}

// GetHourRoomInfoDto() 从对象池中获取HourRoomInfoDto
func GetHourRoomInfoDto() *HourRoomInfoDto {
	return poolHourRoomInfoDto.Get().(*HourRoomInfoDto)
}

// ReleaseHourRoomInfoDto 释放HourRoomInfoDto
func ReleaseHourRoomInfoDto(v *HourRoomInfoDto) {
	v.HourRoomCheckInStart = ""
	v.HourRoomCheckInEnd = ""
	v.Hourage = 0
	poolHourRoomInfoDto.Put(v)
}
