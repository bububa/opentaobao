package alitripmerchant

import (
	"sync"
)

// HotelPictureDto 结构体
type HotelPictureDto struct {
	// 图片集合
	PictureAddress []string `json:"picture_address,omitempty" xml:"picture_address>string,omitempty"`
	// 类型名称
	TypeName string `json:"type_name,omitempty" xml:"type_name,omitempty"`
	// 类型编码
	TypeCode string `json:"type_code,omitempty" xml:"type_code,omitempty"`
}

var poolHotelPictureDto = sync.Pool{
	New: func() any {
		return new(HotelPictureDto)
	},
}

// GetHotelPictureDto() 从对象池中获取HotelPictureDto
func GetHotelPictureDto() *HotelPictureDto {
	return poolHotelPictureDto.Get().(*HotelPictureDto)
}

// ReleaseHotelPictureDto 释放HotelPictureDto
func ReleaseHotelPictureDto(v *HotelPictureDto) {
	v.PictureAddress = v.PictureAddress[:0]
	v.TypeName = ""
	v.TypeCode = ""
	poolHotelPictureDto.Put(v)
}
