package xhotelonlineorder

import (
	"sync"
)

// DailyRoomFee 结构体
type DailyRoomFee struct {
	// 无
	DailyPrices []DailyPrice `json:"daily_prices,omitempty" xml:"daily_prices>daily_price,omitempty"`
}

var poolDailyRoomFee = sync.Pool{
	New: func() any {
		return new(DailyRoomFee)
	},
}

// GetDailyRoomFee() 从对象池中获取DailyRoomFee
func GetDailyRoomFee() *DailyRoomFee {
	return poolDailyRoomFee.Get().(*DailyRoomFee)
}

// ReleaseDailyRoomFee 释放DailyRoomFee
func ReleaseDailyRoomFee(v *DailyRoomFee) {
	v.DailyPrices = v.DailyPrices[:0]
	poolDailyRoomFee.Put(v)
}
