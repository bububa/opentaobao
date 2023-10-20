package xhotelitem

import (
	"sync"
)

// DailyBookingInfo 结构体
type DailyBookingInfo struct {
	// 生效星期，星期一星期二生效就填1，2
	ValidWeeks []string `json:"valid_weeks,omitempty" xml:"valid_weeks>string,omitempty"`
	// 失效时间
	InvalidDates []InvalidDate `json:"invalid_dates,omitempty" xml:"invalid_dates>invalid_date,omitempty"`
	// 入住开始
	CheckInFrom string `json:"check_in_from,omitempty" xml:"check_in_from,omitempty"`
	// 入住结束
	CheckInTo string `json:"check_in_to,omitempty" xml:"check_in_to,omitempty"`
	// 活动折扣
	InvestmentNumber int64 `json:"investment_number,omitempty" xml:"investment_number,omitempty"`
}

var poolDailyBookingInfo = sync.Pool{
	New: func() any {
		return new(DailyBookingInfo)
	},
}

// GetDailyBookingInfo() 从对象池中获取DailyBookingInfo
func GetDailyBookingInfo() *DailyBookingInfo {
	return poolDailyBookingInfo.Get().(*DailyBookingInfo)
}

// ReleaseDailyBookingInfo 释放DailyBookingInfo
func ReleaseDailyBookingInfo(v *DailyBookingInfo) {
	v.ValidWeeks = v.ValidWeeks[:0]
	v.InvalidDates = v.InvalidDates[:0]
	v.CheckInFrom = ""
	v.CheckInTo = ""
	v.InvestmentNumber = 0
	poolDailyBookingInfo.Put(v)
}
