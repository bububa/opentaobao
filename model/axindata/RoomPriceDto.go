package axindata

import (
	"sync"
)

// RoomPriceDto 结构体
type RoomPriceDto struct {
	// rate信息列表
	RateList []RateDto `json:"rate_list,omitempty" xml:"rate_list>rate_dto,omitempty"`
	// 房型信息
	StdRoomInfo *StdRoomType `json:"std_room_info,omitempty" xml:"std_room_info,omitempty"`
}

var poolRoomPriceDto = sync.Pool{
	New: func() any {
		return new(RoomPriceDto)
	},
}

// GetRoomPriceDto() 从对象池中获取RoomPriceDto
func GetRoomPriceDto() *RoomPriceDto {
	return poolRoomPriceDto.Get().(*RoomPriceDto)
}

// ReleaseRoomPriceDto 释放RoomPriceDto
func ReleaseRoomPriceDto(v *RoomPriceDto) {
	v.RateList = v.RateList[:0]
	v.StdRoomInfo = nil
	poolRoomPriceDto.Put(v)
}
