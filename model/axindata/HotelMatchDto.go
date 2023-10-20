package axindata

import (
	"sync"
)

// HotelMatchDto 结构体
type HotelMatchDto struct {
	// 酒店地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 酒店英文名称
	HotelEnName string `json:"hotel_en_name,omitempty" xml:"hotel_en_name,omitempty"`
	// 经度
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 酒店中文名称
	HotelName string `json:"hotel_name,omitempty" xml:"hotel_name,omitempty"`
	// 英文地址
	AddressEn string `json:"address_en,omitempty" xml:"address_en,omitempty"`
	// 纬度
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 酒店电话
	Tel string `json:"tel,omitempty" xml:"tel,omitempty"`
}

var poolHotelMatchDto = sync.Pool{
	New: func() any {
		return new(HotelMatchDto)
	},
}

// GetHotelMatchDto() 从对象池中获取HotelMatchDto
func GetHotelMatchDto() *HotelMatchDto {
	return poolHotelMatchDto.Get().(*HotelMatchDto)
}

// ReleaseHotelMatchDto 释放HotelMatchDto
func ReleaseHotelMatchDto(v *HotelMatchDto) {
	v.Address = ""
	v.HotelEnName = ""
	v.Latitude = ""
	v.HotelName = ""
	v.AddressEn = ""
	v.Longitude = ""
	v.Tel = ""
	poolHotelMatchDto.Put(v)
}
