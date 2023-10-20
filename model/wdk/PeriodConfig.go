package wdk

import (
	"sync"
)

// PeriodConfig 结构体
type PeriodConfig struct {
	// 每天的那些时间段生效
	EveryDayPeriods []string `json:"every_day_periods,omitempty" xml:"every_day_periods>string,omitempty"`
	// 一周的哪几天生效
	Weekdays []string `json:"weekdays,omitempty" xml:"weekdays>string,omitempty"`
}

var poolPeriodConfig = sync.Pool{
	New: func() any {
		return new(PeriodConfig)
	},
}

// GetPeriodConfig() 从对象池中获取PeriodConfig
func GetPeriodConfig() *PeriodConfig {
	return poolPeriodConfig.Get().(*PeriodConfig)
}

// ReleasePeriodConfig 释放PeriodConfig
func ReleasePeriodConfig(v *PeriodConfig) {
	v.EveryDayPeriods = v.EveryDayPeriods[:0]
	v.Weekdays = v.Weekdays[:0]
	poolPeriodConfig.Put(v)
}
