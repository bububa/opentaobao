package xhotelonlineorder

import (
	"sync"
)

// TopDailyInfo 结构体
type TopDailyInfo struct {
	// 日期
	Day string `json:"day,omitempty" xml:"day,omitempty"`
	// 价格
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 早餐数
	BreakFastCount int64 `json:"break_fast_count,omitempty" xml:"break_fast_count,omitempty"`
	// 加价金额
	ExtraAddPrice int64 `json:"extra_add_price,omitempty" xml:"extra_add_price,omitempty"`
}

var poolTopDailyInfo = sync.Pool{
	New: func() any {
		return new(TopDailyInfo)
	},
}

// GetTopDailyInfo() 从对象池中获取TopDailyInfo
func GetTopDailyInfo() *TopDailyInfo {
	return poolTopDailyInfo.Get().(*TopDailyInfo)
}

// ReleaseTopDailyInfo 释放TopDailyInfo
func ReleaseTopDailyInfo(v *TopDailyInfo) {
	v.Day = ""
	v.Price = 0
	v.BreakFastCount = 0
	v.ExtraAddPrice = 0
	poolTopDailyInfo.Put(v)
}
