package travel

import (
	"sync"
)

// ItemHotelInfo 结构体
type ItemHotelInfo struct {
	// 结构化酒店信息 酒店结构化信息和文本描述二选一
	HotelList []ItemHotel `json:"hotel_list,omitempty" xml:"hotel_list>item_hotel,omitempty"`
}

var poolItemHotelInfo = sync.Pool{
	New: func() any {
		return new(ItemHotelInfo)
	},
}

// GetItemHotelInfo() 从对象池中获取ItemHotelInfo
func GetItemHotelInfo() *ItemHotelInfo {
	return poolItemHotelInfo.Get().(*ItemHotelInfo)
}

// ReleaseItemHotelInfo 释放ItemHotelInfo
func ReleaseItemHotelInfo(v *ItemHotelInfo) {
	v.HotelList = v.HotelList[:0]
	poolItemHotelInfo.Put(v)
}
