package xhotel

// PromoInfo 结构体
type PromoInfo struct {
	// 连住优惠的入参
	LongOrderInfo *LongOrderInfo `json:"long_order_info,omitempty" xml:"long_order_info,omitempty"`
	// 早定优惠的入参
	EarlyBookingInfo *EarlyBookingInfo `json:"early_booking_info,omitempty" xml:"early_booking_info,omitempty"`
	// 天天特惠的入参
	DailyBookingInfo *DailyBookingInfo `json:"daily_booking_info,omitempty" xml:"daily_booking_info,omitempty"`
}
