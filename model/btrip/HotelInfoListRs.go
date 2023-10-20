package btrip

import (
	"sync"
)

// HotelInfoListRs 结构体
type HotelInfoListRs struct {
	// 基础酒店数据列表
	Hotels []HotelDto `json:"hotels,omitempty" xml:"hotels>hotel_dto,omitempty"`
	// 酒店数量
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
}

var poolHotelInfoListRs = sync.Pool{
	New: func() any {
		return new(HotelInfoListRs)
	},
}

// GetHotelInfoListRs() 从对象池中获取HotelInfoListRs
func GetHotelInfoListRs() *HotelInfoListRs {
	return poolHotelInfoListRs.Get().(*HotelInfoListRs)
}

// ReleaseHotelInfoListRs 释放HotelInfoListRs
func ReleaseHotelInfoListRs(v *HotelInfoListRs) {
	v.Hotels = v.Hotels[:0]
	v.Total = 0
	poolHotelInfoListRs.Put(v)
}
