package alitripmerchant

import (
	"sync"
)

// DailyMarkupPrice 结构体
type DailyMarkupPrice struct {
	// 实际加价金额
	DailyPrice string `json:"daily_price,omitempty" xml:"daily_price,omitempty"`
	// 加价日期
	Date string `json:"date,omitempty" xml:"date,omitempty"`
}

var poolDailyMarkupPrice = sync.Pool{
	New: func() any {
		return new(DailyMarkupPrice)
	},
}

// GetDailyMarkupPrice() 从对象池中获取DailyMarkupPrice
func GetDailyMarkupPrice() *DailyMarkupPrice {
	return poolDailyMarkupPrice.Get().(*DailyMarkupPrice)
}

// ReleaseDailyMarkupPrice 释放DailyMarkupPrice
func ReleaseDailyMarkupPrice(v *DailyMarkupPrice) {
	v.DailyPrice = ""
	v.Date = ""
	poolDailyMarkupPrice.Put(v)
}
