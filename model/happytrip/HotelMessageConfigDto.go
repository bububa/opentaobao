package happytrip

import (
	"sync"
)

// HotelMessageConfigDto 结构体
type HotelMessageConfigDto struct {
	// 配置信息
	Segments []string `json:"segments,omitempty" xml:"segments>string,omitempty"`
}

var poolHotelMessageConfigDto = sync.Pool{
	New: func() any {
		return new(HotelMessageConfigDto)
	},
}

// GetHotelMessageConfigDto() 从对象池中获取HotelMessageConfigDto
func GetHotelMessageConfigDto() *HotelMessageConfigDto {
	return poolHotelMessageConfigDto.Get().(*HotelMessageConfigDto)
}

// ReleaseHotelMessageConfigDto 释放HotelMessageConfigDto
func ReleaseHotelMessageConfigDto(v *HotelMessageConfigDto) {
	v.Segments = v.Segments[:0]
	poolHotelMessageConfigDto.Put(v)
}
