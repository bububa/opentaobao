package wdk

import (
	"sync"
)

// PeriodicConfigDto 结构体
type PeriodicConfigDto struct {
	// 每天的什么时间阶段搞活动,精确到秒单位 例如:03:00:00_05:00:00
	EveryDayPeriods []string `json:"every_day_periods,omitempty" xml:"every_day_periods>string,omitempty"`
	// 星期几搞活动 [1:Mon;2:Tues;3:Wed;4:Thur;5:Fri;6:Sat;7:Sun]
	Weekdays []string `json:"weekdays,omitempty" xml:"weekdays>string,omitempty"`
	// 周期配置是否生效
	Periodic bool `json:"periodic,omitempty" xml:"periodic,omitempty"`
}

var poolPeriodicConfigDto = sync.Pool{
	New: func() any {
		return new(PeriodicConfigDto)
	},
}

// GetPeriodicConfigDto() 从对象池中获取PeriodicConfigDto
func GetPeriodicConfigDto() *PeriodicConfigDto {
	return poolPeriodicConfigDto.Get().(*PeriodicConfigDto)
}

// ReleasePeriodicConfigDto 释放PeriodicConfigDto
func ReleasePeriodicConfigDto(v *PeriodicConfigDto) {
	v.EveryDayPeriods = v.EveryDayPeriods[:0]
	v.Weekdays = v.Weekdays[:0]
	v.Periodic = false
	poolPeriodicConfigDto.Put(v)
}
