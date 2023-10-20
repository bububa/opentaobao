package axintrade

import (
	"sync"
)

// BookingPolicyDto 结构体
type BookingPolicyDto struct {
	// 提前预订时间
	AdvanceBookingTime string `json:"advance_booking_time,omitempty" xml:"advance_booking_time,omitempty"`
	// 提前预定天数
	AdvanceBookingDays int64 `json:"advance_booking_days,omitempty" xml:"advance_booking_days,omitempty"`
	// 出票后多少小时可用
	ValidHourLimit int64 `json:"valid_hour_limit,omitempty" xml:"valid_hour_limit,omitempty"`
	// 出票后多少分钟可用
	ValidMinuteLimit int64 `json:"valid_minute_limit,omitempty" xml:"valid_minute_limit,omitempty"`
	// 是否需要选择指定日期
	IsSpecifyTime bool `json:"is_specify_time,omitempty" xml:"is_specify_time,omitempty"`
}

var poolBookingPolicyDto = sync.Pool{
	New: func() any {
		return new(BookingPolicyDto)
	},
}

// GetBookingPolicyDto() 从对象池中获取BookingPolicyDto
func GetBookingPolicyDto() *BookingPolicyDto {
	return poolBookingPolicyDto.Get().(*BookingPolicyDto)
}

// ReleaseBookingPolicyDto 释放BookingPolicyDto
func ReleaseBookingPolicyDto(v *BookingPolicyDto) {
	v.AdvanceBookingTime = ""
	v.AdvanceBookingDays = 0
	v.ValidHourLimit = 0
	v.ValidMinuteLimit = 0
	v.IsSpecifyTime = false
	poolBookingPolicyDto.Put(v)
}
