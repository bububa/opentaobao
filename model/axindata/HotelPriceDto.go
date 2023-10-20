package axindata

import (
	"sync"
)

// HotelPriceDto 结构体
type HotelPriceDto struct {
	// 房型信息列表
	RoomList []RoomPriceDto `json:"room_list,omitempty" xml:"room_list>room_price_dto,omitempty"`
	// 标准酒店id
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
}

var poolHotelPriceDto = sync.Pool{
	New: func() any {
		return new(HotelPriceDto)
	},
}

// GetHotelPriceDto() 从对象池中获取HotelPriceDto
func GetHotelPriceDto() *HotelPriceDto {
	return poolHotelPriceDto.Get().(*HotelPriceDto)
}

// ReleaseHotelPriceDto 释放HotelPriceDto
func ReleaseHotelPriceDto(v *HotelPriceDto) {
	v.RoomList = v.RoomList[:0]
	v.Shid = 0
	poolHotelPriceDto.Put(v)
}
