package tbtrade

import (
	"sync"
)

// SortInfo 结构体
type SortInfo struct {
	// daySortVal
	DaySortVal int64 `json:"day_sort_val,omitempty" xml:"day_sort_val,omitempty"`
	// hourSortVal
	HourSortVal int64 `json:"hour_sort_val,omitempty" xml:"hour_sort_val,omitempty"`
}

var poolSortInfo = sync.Pool{
	New: func() any {
		return new(SortInfo)
	},
}

// GetSortInfo() 从对象池中获取SortInfo
func GetSortInfo() *SortInfo {
	return poolSortInfo.Get().(*SortInfo)
}

// ReleaseSortInfo 释放SortInfo
func ReleaseSortInfo(v *SortInfo) {
	v.DaySortVal = 0
	v.HourSortVal = 0
	poolSortInfo.Put(v)
}
