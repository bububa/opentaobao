package axintrade

import (
	"sync"
)

// HotelInfo 结构体
type HotelInfo struct {
	// 酒店联系电话
	HotelTel string `json:"hotel_tel,omitempty" xml:"hotel_tel,omitempty"`
	// 酒店地址
	HotelAddress string `json:"hotel_address,omitempty" xml:"hotel_address,omitempty"`
	// 酒店名称
	HotelName string `json:"hotel_name,omitempty" xml:"hotel_name,omitempty"`
	// 酒店id
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
}

var poolHotelInfo = sync.Pool{
	New: func() any {
		return new(HotelInfo)
	},
}

// GetHotelInfo() 从对象池中获取HotelInfo
func GetHotelInfo() *HotelInfo {
	return poolHotelInfo.Get().(*HotelInfo)
}

// ReleaseHotelInfo 释放HotelInfo
func ReleaseHotelInfo(v *HotelInfo) {
	v.HotelTel = ""
	v.HotelAddress = ""
	v.HotelName = ""
	v.Shid = 0
	poolHotelInfo.Put(v)
}
