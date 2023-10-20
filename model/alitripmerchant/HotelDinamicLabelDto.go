package alitripmerchant

import (
	"sync"
)

// HotelDinamicLabelDto 结构体
type HotelDinamicLabelDto struct {
	// 权益名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 权益类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
}

var poolHotelDinamicLabelDto = sync.Pool{
	New: func() any {
		return new(HotelDinamicLabelDto)
	},
}

// GetHotelDinamicLabelDto() 从对象池中获取HotelDinamicLabelDto
func GetHotelDinamicLabelDto() *HotelDinamicLabelDto {
	return poolHotelDinamicLabelDto.Get().(*HotelDinamicLabelDto)
}

// ReleaseHotelDinamicLabelDto 释放HotelDinamicLabelDto
func ReleaseHotelDinamicLabelDto(v *HotelDinamicLabelDto) {
	v.Name = ""
	v.Type = ""
	poolHotelDinamicLabelDto.Put(v)
}
