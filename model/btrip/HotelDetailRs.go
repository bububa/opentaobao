package btrip

import (
	"sync"
)

// HotelDetailRs 结构体
type HotelDetailRs struct {
	// 房型列表
	Rooms []HotelDetailRoomDto `json:"rooms,omitempty" xml:"rooms>hotel_detail_room_dto,omitempty"`
	// 入住时间
	CheckIn string `json:"check_in,omitempty" xml:"check_in,omitempty"`
	// 离店时间
	CheckOut string `json:"check_out,omitempty" xml:"check_out,omitempty"`
	// 跟踪ID，可忽略
	EagleTraceId string `json:"eagle_trace_id,omitempty" xml:"eagle_trace_id,omitempty"`
	// 跟踪ID。可忽略
	SearchId string `json:"search_id,omitempty" xml:"search_id,omitempty"`
	// 酒店标准ID
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
}

var poolHotelDetailRs = sync.Pool{
	New: func() any {
		return new(HotelDetailRs)
	},
}

// GetHotelDetailRs() 从对象池中获取HotelDetailRs
func GetHotelDetailRs() *HotelDetailRs {
	return poolHotelDetailRs.Get().(*HotelDetailRs)
}

// ReleaseHotelDetailRs 释放HotelDetailRs
func ReleaseHotelDetailRs(v *HotelDetailRs) {
	v.Rooms = v.Rooms[:0]
	v.CheckIn = ""
	v.CheckOut = ""
	v.EagleTraceId = ""
	v.SearchId = ""
	v.Shid = 0
	poolHotelDetailRs.Put(v)
}
