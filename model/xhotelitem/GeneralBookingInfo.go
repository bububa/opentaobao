package xhotelitem

import (
	"sync"
)

// GeneralBookingInfo 结构体
type GeneralBookingInfo struct {
	// 活动入住开始时间
	CheckInFrom string `json:"check_in_from,omitempty" xml:"check_in_from,omitempty"`
	// 活动离店结束时间
	CheckOutTo string `json:"check_out_to,omitempty" xml:"check_out_to,omitempty"`
	// 活动折扣
	InvestmentNumber int64 `json:"investment_number,omitempty" xml:"investment_number,omitempty"`
}

var poolGeneralBookingInfo = sync.Pool{
	New: func() any {
		return new(GeneralBookingInfo)
	},
}

// GetGeneralBookingInfo() 从对象池中获取GeneralBookingInfo
func GetGeneralBookingInfo() *GeneralBookingInfo {
	return poolGeneralBookingInfo.Get().(*GeneralBookingInfo)
}

// ReleaseGeneralBookingInfo 释放GeneralBookingInfo
func ReleaseGeneralBookingInfo(v *GeneralBookingInfo) {
	v.CheckInFrom = ""
	v.CheckOutTo = ""
	v.InvestmentNumber = 0
	poolGeneralBookingInfo.Put(v)
}
