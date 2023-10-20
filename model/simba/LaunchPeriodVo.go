package simba

import (
	"sync"
)

// LaunchPeriodVo 结构体
type LaunchPeriodVo struct {
	// 第x天各时段的折扣情况
	TimeSpanList []TimeSpanQueryResVo `json:"time_span_list,omitempty" xml:"time_span_list>time_span_query_res_vo,omitempty"`
	// 本周的第x天
	DayOfWeek int64 `json:"day_of_week,omitempty" xml:"day_of_week,omitempty"`
}

var poolLaunchPeriodVo = sync.Pool{
	New: func() any {
		return new(LaunchPeriodVo)
	},
}

// GetLaunchPeriodVo() 从对象池中获取LaunchPeriodVo
func GetLaunchPeriodVo() *LaunchPeriodVo {
	return poolLaunchPeriodVo.Get().(*LaunchPeriodVo)
}

// ReleaseLaunchPeriodVo 释放LaunchPeriodVo
func ReleaseLaunchPeriodVo(v *LaunchPeriodVo) {
	v.TimeSpanList = v.TimeSpanList[:0]
	v.DayOfWeek = 0
	poolLaunchPeriodVo.Put(v)
}
