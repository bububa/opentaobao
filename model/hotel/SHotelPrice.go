package hotel

import (
	"sync"
)

// SHotelPrice 结构体
type SHotelPrice struct {
	// 每个标准酒店某一天的所有库价集合
	DailyPriceList []SHotelDailyPrice `json:"daily_price_list,omitempty" xml:"daily_price_list>s_hotel_daily_price,omitempty"`
	// 标准房型shid
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
}

var poolSHotelPrice = sync.Pool{
	New: func() any {
		return new(SHotelPrice)
	},
}

// GetSHotelPrice() 从对象池中获取SHotelPrice
func GetSHotelPrice() *SHotelPrice {
	return poolSHotelPrice.Get().(*SHotelPrice)
}

// ReleaseSHotelPrice 释放SHotelPrice
func ReleaseSHotelPrice(v *SHotelPrice) {
	v.DailyPriceList = v.DailyPriceList[:0]
	v.Shid = 0
	poolSHotelPrice.Put(v)
}
