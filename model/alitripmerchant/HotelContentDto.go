package alitripmerchant

import (
	"sync"

	"github.com/bububa/opentaobao/model"
)

// HotelContentDto 结构体
type HotelContentDto struct {
	// 酒店id
	HotelId string `json:"hotel_id,omitempty" xml:"hotel_id,omitempty"`
	// 酒店信息
	HotelIntro string `json:"hotel_intro,omitempty" xml:"hotel_intro,omitempty"`
	// 酒店电话
	HotelPhone string `json:"hotel_phone,omitempty" xml:"hotel_phone,omitempty"`
	// 图片
	PictureUrl string `json:"picture_url,omitempty" xml:"picture_url,omitempty"`
	// 纬度
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 经度
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 酒店名称
	HotelName string `json:"hotel_name,omitempty" xml:"hotel_name,omitempty"`
	// 城市名称
	CityCn string `json:"city_cn,omitempty" xml:"city_cn,omitempty"`
	// 城市名称首字母
	CityPyHead string `json:"city_py_head,omitempty" xml:"city_py_head,omitempty"`
	// 地图
	PositionType *model.File `json:"position_type,omitempty" xml:"position_type,omitempty"`
	// 标准id
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
	// 城市code
	CityCode int64 `json:"city_code,omitempty" xml:"city_code,omitempty"`
}

var poolHotelContentDto = sync.Pool{
	New: func() any {
		return new(HotelContentDto)
	},
}

// GetHotelContentDto() 从对象池中获取HotelContentDto
func GetHotelContentDto() *HotelContentDto {
	return poolHotelContentDto.Get().(*HotelContentDto)
}

// ReleaseHotelContentDto 释放HotelContentDto
func ReleaseHotelContentDto(v *HotelContentDto) {
	v.HotelId = ""
	v.HotelIntro = ""
	v.HotelPhone = ""
	v.PictureUrl = ""
	v.Latitude = ""
	v.Longitude = ""
	v.Address = ""
	v.HotelName = ""
	v.CityCn = ""
	v.CityPyHead = ""
	v.PositionType = nil
	v.Shid = 0
	v.CityCode = 0
	poolHotelContentDto.Put(v)
}
