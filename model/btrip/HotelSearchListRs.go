package btrip

import (
	"sync"
)

// HotelSearchListRs 结构体
type HotelSearchListRs struct {
	// 酒店列表
	Hotels []HotelListDto `json:"hotels,omitempty" xml:"hotels>hotel_list_dto,omitempty"`
	// 酒店数量
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
}

var poolHotelSearchListRs = sync.Pool{
	New: func() any {
		return new(HotelSearchListRs)
	},
}

// GetHotelSearchListRs() 从对象池中获取HotelSearchListRs
func GetHotelSearchListRs() *HotelSearchListRs {
	return poolHotelSearchListRs.Get().(*HotelSearchListRs)
}

// ReleaseHotelSearchListRs 释放HotelSearchListRs
func ReleaseHotelSearchListRs(v *HotelSearchListRs) {
	v.Hotels = v.Hotels[:0]
	v.Total = 0
	poolHotelSearchListRs.Put(v)
}
