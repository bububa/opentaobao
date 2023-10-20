package btrip

import (
	"sync"
)

// HotHotelIdListRs 结构体
type HotHotelIdListRs struct {
	// 酒店Id列表
	HotelIds []string `json:"hotel_ids,omitempty" xml:"hotel_ids>string,omitempty"`
}

var poolHotHotelIdListRs = sync.Pool{
	New: func() any {
		return new(HotHotelIdListRs)
	},
}

// GetHotHotelIdListRs() 从对象池中获取HotHotelIdListRs
func GetHotHotelIdListRs() *HotHotelIdListRs {
	return poolHotHotelIdListRs.Get().(*HotHotelIdListRs)
}

// ReleaseHotHotelIdListRs 释放HotHotelIdListRs
func ReleaseHotHotelIdListRs(v *HotHotelIdListRs) {
	v.HotelIds = v.HotelIds[:0]
	poolHotHotelIdListRs.Put(v)
}
