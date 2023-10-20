package axindata

import (
	"sync"
)

// HotelRoomMatchDto 结构体
type HotelRoomMatchDto struct {
	// 房型英文名称
	RoomNameEn string `json:"room_name_en,omitempty" xml:"room_name_en,omitempty"`
	// 房型中文名称
	RoomName string `json:"room_name,omitempty" xml:"room_name,omitempty"`
	// 床型类型
	BedType string `json:"bed_type,omitempty" xml:"bed_type,omitempty"`
	// 对应标准酒店id
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
}

var poolHotelRoomMatchDto = sync.Pool{
	New: func() any {
		return new(HotelRoomMatchDto)
	},
}

// GetHotelRoomMatchDto() 从对象池中获取HotelRoomMatchDto
func GetHotelRoomMatchDto() *HotelRoomMatchDto {
	return poolHotelRoomMatchDto.Get().(*HotelRoomMatchDto)
}

// ReleaseHotelRoomMatchDto 释放HotelRoomMatchDto
func ReleaseHotelRoomMatchDto(v *HotelRoomMatchDto) {
	v.RoomNameEn = ""
	v.RoomName = ""
	v.BedType = ""
	v.Shid = 0
	poolHotelRoomMatchDto.Put(v)
}
