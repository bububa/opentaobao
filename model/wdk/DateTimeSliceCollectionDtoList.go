package wdk

import (
	"sync"
)

// DateTimeSliceCollectionDtoList 结构体
type DateTimeSliceCollectionDtoList struct {
	// 时间片
	TimeSliceList []TimeSliceList `json:"time_slice_list,omitempty" xml:"time_slice_list>time_slice_list,omitempty"`
	// 日期
	Date string `json:"date,omitempty" xml:"date,omitempty"`
}

var poolDateTimeSliceCollectionDtoList = sync.Pool{
	New: func() any {
		return new(DateTimeSliceCollectionDtoList)
	},
}

// GetDateTimeSliceCollectionDtoList() 从对象池中获取DateTimeSliceCollectionDtoList
func GetDateTimeSliceCollectionDtoList() *DateTimeSliceCollectionDtoList {
	return poolDateTimeSliceCollectionDtoList.Get().(*DateTimeSliceCollectionDtoList)
}

// ReleaseDateTimeSliceCollectionDtoList 释放DateTimeSliceCollectionDtoList
func ReleaseDateTimeSliceCollectionDtoList(v *DateTimeSliceCollectionDtoList) {
	v.TimeSliceList = v.TimeSliceList[:0]
	v.Date = ""
	poolDateTimeSliceCollectionDtoList.Put(v)
}
