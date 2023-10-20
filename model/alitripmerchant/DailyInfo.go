package alitripmerchant

import (
	"sync"
)

// DailyInfo 结构体
type DailyInfo struct {
	// 日期
	Day string `json:"day,omitempty" xml:"day,omitempty"`
	// 日期字符串格式
	DayStr string `json:"day_str,omitempty" xml:"day_str,omitempty"`
	// 早餐名称
	BreakfastName string `json:"breakfast_name,omitempty" xml:"breakfast_name,omitempty"`
	// 不含税价格
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 每天的税，单位分
	Tax int64 `json:"tax,omitempty" xml:"tax,omitempty"`
	// 早餐数量
	BreakfastCount int64 `json:"breakfast_count,omitempty" xml:"breakfast_count,omitempty"`
	// 外币不含税价格
	OutPrice int64 `json:"out_price,omitempty" xml:"out_price,omitempty"`
	// 外币每天的税，单位分
	OutTax int64 `json:"out_tax,omitempty" xml:"out_tax,omitempty"`
}

var poolDailyInfo = sync.Pool{
	New: func() any {
		return new(DailyInfo)
	},
}

// GetDailyInfo() 从对象池中获取DailyInfo
func GetDailyInfo() *DailyInfo {
	return poolDailyInfo.Get().(*DailyInfo)
}

// ReleaseDailyInfo 释放DailyInfo
func ReleaseDailyInfo(v *DailyInfo) {
	v.Day = ""
	v.DayStr = ""
	v.BreakfastName = ""
	v.Price = 0
	v.Tax = 0
	v.BreakfastCount = 0
	v.OutPrice = 0
	v.OutTax = 0
	poolDailyInfo.Put(v)
}
