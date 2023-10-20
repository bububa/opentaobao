package xhotelonlineorder

import (
	"sync"
)

// DailyPriceInfo 结构体
type DailyPriceInfo struct {
	// 日历日期
	Day string `json:"day,omitempty" xml:"day,omitempty"`
	// 日历早餐
	BreakFastCount int64 `json:"break_fast_count,omitempty" xml:"break_fast_count,omitempty"`
	// 日历价格
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 如果是低价加价商品，此价格是底价。如果是非底价商品且为会员商品，则为会员结算价
	BasePrice int64 `json:"base_price,omitempty" xml:"base_price,omitempty"`
}

var poolDailyPriceInfo = sync.Pool{
	New: func() any {
		return new(DailyPriceInfo)
	},
}

// GetDailyPriceInfo() 从对象池中获取DailyPriceInfo
func GetDailyPriceInfo() *DailyPriceInfo {
	return poolDailyPriceInfo.Get().(*DailyPriceInfo)
}

// ReleaseDailyPriceInfo 释放DailyPriceInfo
func ReleaseDailyPriceInfo(v *DailyPriceInfo) {
	v.Day = ""
	v.BreakFastCount = 0
	v.Price = 0
	v.BasePrice = 0
	poolDailyPriceInfo.Put(v)
}
