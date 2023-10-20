package axintrade

import (
	"sync"
)

// HotelRateDto 结构体
type HotelRateDto struct {
	// 酒店名称
	HotelName string `json:"hotel_name,omitempty" xml:"hotel_name,omitempty"`
	// 房型名称
	RoomName string `json:"room_name,omitempty" xml:"room_name,omitempty"`
}

var poolHotelRateDto = sync.Pool{
	New: func() any {
		return new(HotelRateDto)
	},
}

// GetHotelRateDto() 从对象池中获取HotelRateDto
func GetHotelRateDto() *HotelRateDto {
	return poolHotelRateDto.Get().(*HotelRateDto)
}

// ReleaseHotelRateDto 释放HotelRateDto
func ReleaseHotelRateDto(v *HotelRateDto) {
	v.HotelName = ""
	v.RoomName = ""
	poolHotelRateDto.Put(v)
}
