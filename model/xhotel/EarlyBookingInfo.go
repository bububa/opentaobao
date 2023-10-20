package xhotel

import (
	"sync"
)

// EarlyBookingInfo 结构体
type EarlyBookingInfo struct {
	// 最少提前预定天数，数字范围限定在1-60
	MinPreBookingDays int64 `json:"min_pre_booking_days,omitempty" xml:"min_pre_booking_days,omitempty"`
	// 折扣比例，填30就意味着原价的30%，也就是打3折。数字范围限定在10-95之间
	InvestmentNumber int64 `json:"investment_number,omitempty" xml:"investment_number,omitempty"`
	// 连住天数,可为空，范围1-60
	MinContinuityStay int64 `json:"min_continuity_stay,omitempty" xml:"min_continuity_stay,omitempty"`
}

var poolEarlyBookingInfo = sync.Pool{
	New: func() any {
		return new(EarlyBookingInfo)
	},
}

// GetEarlyBookingInfo() 从对象池中获取EarlyBookingInfo
func GetEarlyBookingInfo() *EarlyBookingInfo {
	return poolEarlyBookingInfo.Get().(*EarlyBookingInfo)
}

// ReleaseEarlyBookingInfo 释放EarlyBookingInfo
func ReleaseEarlyBookingInfo(v *EarlyBookingInfo) {
	v.MinPreBookingDays = 0
	v.InvestmentNumber = 0
	v.MinContinuityStay = 0
	poolEarlyBookingInfo.Put(v)
}
