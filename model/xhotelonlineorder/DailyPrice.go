package xhotelonlineorder

import (
	"sync"
)

// DailyPrice 结构体
type DailyPrice struct {
	// 日期
	Date string `json:"date,omitempty" xml:"date,omitempty"`
	// 结账状态(1:结账,0:未结账)
	Checkout int64 `json:"checkout,omitempty" xml:"checkout,omitempty"`
	// 每日实际房费
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
}

var poolDailyPrice = sync.Pool{
	New: func() any {
		return new(DailyPrice)
	},
}

// GetDailyPrice() 从对象池中获取DailyPrice
func GetDailyPrice() *DailyPrice {
	return poolDailyPrice.Get().(*DailyPrice)
}

// ReleaseDailyPrice 释放DailyPrice
func ReleaseDailyPrice(v *DailyPrice) {
	v.Date = ""
	v.Checkout = 0
	v.Price = 0
	poolDailyPrice.Put(v)
}
