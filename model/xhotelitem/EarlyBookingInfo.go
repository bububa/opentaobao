package xhotelitem

// EarlyBookingInfo 结构体
type EarlyBookingInfo struct {
	// 活动折扣
	InvestmentNumber int64 `json:"investment_number,omitempty" xml:"investment_number,omitempty"`
	// 早定天数
	MinPreBookingDays int64 `json:"min_pre_booking_days,omitempty" xml:"min_pre_booking_days,omitempty"`
	// 连住天数
	MinContinuityStay int64 `json:"min_continuity_stay,omitempty" xml:"min_continuity_stay,omitempty"`
}
