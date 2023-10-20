package alitripmerchant

import (
	"sync"
)

// HotelAddressDetail 结构体
type HotelAddressDetail struct {
	// 酒店id
	HotelId string `json:"hotel_id,omitempty" xml:"hotel_id,omitempty"`
	// 酒店头图url
	PictureUrl string `json:"picture_url,omitempty" xml:"picture_url,omitempty"`
	// address
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 酒店中文名称
	NameCn string `json:"name_cn,omitempty" xml:"name_cn,omitempty"`
	// 标准库id
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
}

var poolHotelAddressDetail = sync.Pool{
	New: func() any {
		return new(HotelAddressDetail)
	},
}

// GetHotelAddressDetail() 从对象池中获取HotelAddressDetail
func GetHotelAddressDetail() *HotelAddressDetail {
	return poolHotelAddressDetail.Get().(*HotelAddressDetail)
}

// ReleaseHotelAddressDetail 释放HotelAddressDetail
func ReleaseHotelAddressDetail(v *HotelAddressDetail) {
	v.HotelId = ""
	v.PictureUrl = ""
	v.Address = ""
	v.NameCn = ""
	v.Shid = 0
	poolHotelAddressDetail.Put(v)
}
