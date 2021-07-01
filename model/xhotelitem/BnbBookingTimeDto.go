package xhotelitem

// BnbBookingTimeDto 结构体
type BnbBookingTimeDto struct {
	// 最晚预定时间 hh:mm,24小时时间格式
	LatestBookingTime string `json:"latest_booking_time,omitempty" xml:"latest_booking_time,omitempty"`
	// 最早入住时间 hh:mm,24小时时间格式住时间
	EarliestCheckInTime string `json:"earliest_check_in_time,omitempty" xml:"earliest_check_in_time,omitempty"`
	// 最晚入住时间 hh:mm,24小时时间格式
	LatestCheckInTime string `json:"latest_check_in_time,omitempty" xml:"latest_check_in_time,omitempty"`
	// 最晚离店时间 hh:mm,24小时时间格式
	LatestCheckOutTime string `json:"latest_check_out_time,omitempty" xml:"latest_check_out_time,omitempty"`
	// 开始接待时间 hh:mm,24小时时间格式
	StartReceptionTime string `json:"start_reception_time,omitempty" xml:"start_reception_time,omitempty"`
	// 结束接待时间 hh:mm,24小时时间格式
	EndReceptionTime string `json:"end_reception_time,omitempty" xml:"end_reception_time,omitempty"`
}
