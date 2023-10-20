package travel

import (
	"sync"
)

// PontusTravelItemHotelInfo 结构体
type PontusTravelItemHotelInfo struct {
	// 结构化酒店信息 酒店结构化信息和文本描述二选一
	HotelList []PontusTravelItemHotel `json:"hotel_list,omitempty" xml:"hotel_list>pontus_travel_item_hotel,omitempty"`
	// 酒店描述文本，&amp;lt;1500字符 酒店结构化信息和文本描述二选一
	HotelDesc string `json:"hotel_desc,omitempty" xml:"hotel_desc,omitempty"`
	// 必填，默认为0 必须大于等于0, 且小于或等于行程晚数
	HotelDays int64 `json:"hotel_days,omitempty" xml:"hotel_days,omitempty"`
}

var poolPontusTravelItemHotelInfo = sync.Pool{
	New: func() any {
		return new(PontusTravelItemHotelInfo)
	},
}

// GetPontusTravelItemHotelInfo() 从对象池中获取PontusTravelItemHotelInfo
func GetPontusTravelItemHotelInfo() *PontusTravelItemHotelInfo {
	return poolPontusTravelItemHotelInfo.Get().(*PontusTravelItemHotelInfo)
}

// ReleasePontusTravelItemHotelInfo 释放PontusTravelItemHotelInfo
func ReleasePontusTravelItemHotelInfo(v *PontusTravelItemHotelInfo) {
	v.HotelList = v.HotelList[:0]
	v.HotelDesc = ""
	v.HotelDays = 0
	poolPontusTravelItemHotelInfo.Put(v)
}
