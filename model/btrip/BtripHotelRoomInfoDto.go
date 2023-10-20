package btrip

import (
	"sync"
)

// BtripHotelRoomInfoDto 结构体
type BtripHotelRoomInfoDto struct {
	// 酒店房间设施列表
	RoomFacilityList []string `json:"room_facility_list,omitempty" xml:"room_facility_list>string,omitempty"`
	// 房型名称
	RoomTypeName string `json:"room_type_name,omitempty" xml:"room_type_name,omitempty"`
}

var poolBtripHotelRoomInfoDto = sync.Pool{
	New: func() any {
		return new(BtripHotelRoomInfoDto)
	},
}

// GetBtripHotelRoomInfoDto() 从对象池中获取BtripHotelRoomInfoDto
func GetBtripHotelRoomInfoDto() *BtripHotelRoomInfoDto {
	return poolBtripHotelRoomInfoDto.Get().(*BtripHotelRoomInfoDto)
}

// ReleaseBtripHotelRoomInfoDto 释放BtripHotelRoomInfoDto
func ReleaseBtripHotelRoomInfoDto(v *BtripHotelRoomInfoDto) {
	v.RoomFacilityList = v.RoomFacilityList[:0]
	v.RoomTypeName = ""
	poolBtripHotelRoomInfoDto.Put(v)
}
