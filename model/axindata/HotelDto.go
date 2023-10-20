package axindata

import (
	"sync"
)

// HotelDto 结构体
type HotelDto struct {
	// 标准酒店id
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
	// 城市code
	CityCode int64 `json:"city_code,omitempty" xml:"city_code,omitempty"`
}

var poolHotelDto = sync.Pool{
	New: func() any {
		return new(HotelDto)
	},
}

// GetHotelDto() 从对象池中获取HotelDto
func GetHotelDto() *HotelDto {
	return poolHotelDto.Get().(*HotelDto)
}

// ReleaseHotelDto 释放HotelDto
func ReleaseHotelDto(v *HotelDto) {
	v.Shid = 0
	v.CityCode = 0
	poolHotelDto.Put(v)
}
