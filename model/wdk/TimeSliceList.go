package wdk

import (
	"sync"
)

// TimeSliceList 结构体
type TimeSliceList struct {
	// 时间片
	TimeSlot string `json:"time_slot,omitempty" xml:"time_slot,omitempty"`
}

var poolTimeSliceList = sync.Pool{
	New: func() any {
		return new(TimeSliceList)
	},
}

// GetTimeSliceList() 从对象池中获取TimeSliceList
func GetTimeSliceList() *TimeSliceList {
	return poolTimeSliceList.Get().(*TimeSliceList)
}

// ReleaseTimeSliceList 释放TimeSliceList
func ReleaseTimeSliceList(v *TimeSliceList) {
	v.TimeSlot = ""
	poolTimeSliceList.Put(v)
}
