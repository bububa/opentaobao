package kclub

import (
	"sync"
)

// QaPvDto 结构体
type QaPvDto struct {
	// 十天访问量
	TenDayPv int64 `json:"ten_day_pv,omitempty" xml:"ten_day_pv,omitempty"`
	// 7天访问量
	SevenDayPv int64 `json:"seven_day_pv,omitempty" xml:"seven_day_pv,omitempty"`
	// 5天访问量
	FiveDayPv int64 `json:"five_day_pv,omitempty" xml:"five_day_pv,omitempty"`
	// 3天访问量
	ThreeDayPv int64 `json:"three_day_pv,omitempty" xml:"three_day_pv,omitempty"`
	// 1天访问量
	OneDayPv int64 `json:"one_day_pv,omitempty" xml:"one_day_pv,omitempty"`
}

var poolQaPvDto = sync.Pool{
	New: func() any {
		return new(QaPvDto)
	},
}

// GetQaPvDto() 从对象池中获取QaPvDto
func GetQaPvDto() *QaPvDto {
	return poolQaPvDto.Get().(*QaPvDto)
}

// ReleaseQaPvDto 释放QaPvDto
func ReleaseQaPvDto(v *QaPvDto) {
	v.TenDayPv = 0
	v.SevenDayPv = 0
	v.FiveDayPv = 0
	v.ThreeDayPv = 0
	v.OneDayPv = 0
	poolQaPvDto.Put(v)
}
