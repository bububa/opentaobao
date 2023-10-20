package ascp

import (
	"sync"
)

// Period 结构体
type Period struct {
	// 开始时间
	BeginTime int64 `json:"begin_time,omitempty" xml:"begin_time,omitempty"`
	// 结束时间
	EndTime int64 `json:"end_time,omitempty" xml:"end_time,omitempty"`
}

var poolPeriod = sync.Pool{
	New: func() any {
		return new(Period)
	},
}

// GetPeriod() 从对象池中获取Period
func GetPeriod() *Period {
	return poolPeriod.Get().(*Period)
}

// ReleasePeriod 释放Period
func ReleasePeriod(v *Period) {
	v.BeginTime = 0
	v.EndTime = 0
	poolPeriod.Put(v)
}
