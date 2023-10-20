package hotelhstdf

import (
	"sync"
)

// AlitripHotelHstdfShotelRoomtypeMappingsListResults 结构体
type AlitripHotelHstdfShotelRoomtypeMappingsListResults struct {
	// 卖家房型
	RoomType *RoomTypePo `json:"room_type,omitempty" xml:"room_type,omitempty"`
	// 标准房型
	StdRoomType *SroomTypePo `json:"std_room_type,omitempty" xml:"std_room_type,omitempty"`
	// 是否匹配
	Matched bool `json:"matched,omitempty" xml:"matched,omitempty"`
}

var poolAlitripHotelHstdfShotelRoomtypeMappingsListResults = sync.Pool{
	New: func() any {
		return new(AlitripHotelHstdfShotelRoomtypeMappingsListResults)
	},
}

// GetAlitripHotelHstdfShotelRoomtypeMappingsListResults() 从对象池中获取AlitripHotelHstdfShotelRoomtypeMappingsListResults
func GetAlitripHotelHstdfShotelRoomtypeMappingsListResults() *AlitripHotelHstdfShotelRoomtypeMappingsListResults {
	return poolAlitripHotelHstdfShotelRoomtypeMappingsListResults.Get().(*AlitripHotelHstdfShotelRoomtypeMappingsListResults)
}

// ReleaseAlitripHotelHstdfShotelRoomtypeMappingsListResults 释放AlitripHotelHstdfShotelRoomtypeMappingsListResults
func ReleaseAlitripHotelHstdfShotelRoomtypeMappingsListResults(v *AlitripHotelHstdfShotelRoomtypeMappingsListResults) {
	v.RoomType = nil
	v.StdRoomType = nil
	v.Matched = false
	poolAlitripHotelHstdfShotelRoomtypeMappingsListResults.Put(v)
}
