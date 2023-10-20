package wdk

import (
	"sync"
)

// PromiseTimeSlice 结构体
type PromiseTimeSlice struct {
	// 时间片小时维度列表, 单个时间片格式为09:00-09:30
	SliceList []string `json:"slice_list,omitempty" xml:"slice_list>string,omitempty"`
	// 时间片日期 YYYY-MM-DD
	SliceDate string `json:"slice_date,omitempty" xml:"slice_date,omitempty"`
}

var poolPromiseTimeSlice = sync.Pool{
	New: func() any {
		return new(PromiseTimeSlice)
	},
}

// GetPromiseTimeSlice() 从对象池中获取PromiseTimeSlice
func GetPromiseTimeSlice() *PromiseTimeSlice {
	return poolPromiseTimeSlice.Get().(*PromiseTimeSlice)
}

// ReleasePromiseTimeSlice 释放PromiseTimeSlice
func ReleasePromiseTimeSlice(v *PromiseTimeSlice) {
	v.SliceList = v.SliceList[:0]
	v.SliceDate = ""
	poolPromiseTimeSlice.Put(v)
}
