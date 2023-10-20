package xhotel

import (
	"sync"
)

// PromoInfo 结构体
type PromoInfo struct {
	// 连住优惠的入参
	LongOrderInfo *LongOrderInfo `json:"long_order_info,omitempty" xml:"long_order_info,omitempty"`
	// 早定优惠的入参
	EarlyBookingInfo *EarlyBookingInfo `json:"early_booking_info,omitempty" xml:"early_booking_info,omitempty"`
	// 天天特惠的入参
	DailyBookingInfo *DailyBookingInfo `json:"daily_booking_info,omitempty" xml:"daily_booking_info,omitempty"`
}

var poolPromoInfo = sync.Pool{
	New: func() any {
		return new(PromoInfo)
	},
}

// GetPromoInfo() 从对象池中获取PromoInfo
func GetPromoInfo() *PromoInfo {
	return poolPromoInfo.Get().(*PromoInfo)
}

// ReleasePromoInfo 释放PromoInfo
func ReleasePromoInfo(v *PromoInfo) {
	v.LongOrderInfo = nil
	v.EarlyBookingInfo = nil
	v.DailyBookingInfo = nil
	poolPromoInfo.Put(v)
}
