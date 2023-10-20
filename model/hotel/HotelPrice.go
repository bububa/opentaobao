package hotel

import (
	"sync"
)

// HotelPrice 结构体
type HotelPrice struct {
	// 房型报价列表
	RoomPrices []RoomPrice `json:"room_prices,omitempty" xml:"room_prices>room_price,omitempty"`
	// 酒店名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 酒店id
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
}

var poolHotelPrice = sync.Pool{
	New: func() any {
		return new(HotelPrice)
	},
}

// GetHotelPrice() 从对象池中获取HotelPrice
func GetHotelPrice() *HotelPrice {
	return poolHotelPrice.Get().(*HotelPrice)
}

// ReleaseHotelPrice 释放HotelPrice
func ReleaseHotelPrice(v *HotelPrice) {
	v.RoomPrices = v.RoomPrices[:0]
	v.Name = ""
	v.Shid = 0
	poolHotelPrice.Put(v)
}
