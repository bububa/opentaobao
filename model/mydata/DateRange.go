package mydata

import (
	"sync"
)

// DateRange 结构体
type DateRange struct {
	// 数据周期结束日期（含）
	EndDate string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// 数据周期开始日期（含）
	StartDate string `json:"start_date,omitempty" xml:"start_date,omitempty"`
}

var poolDateRange = sync.Pool{
	New: func() any {
		return new(DateRange)
	},
}

// GetDateRange() 从对象池中获取DateRange
func GetDateRange() *DateRange {
	return poolDateRange.Get().(*DateRange)
}

// ReleaseDateRange 释放DateRange
func ReleaseDateRange(v *DateRange) {
	v.EndDate = ""
	v.StartDate = ""
	poolDateRange.Put(v)
}
