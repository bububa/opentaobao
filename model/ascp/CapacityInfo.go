package ascp

import (
	"sync"
)

// CapacityInfo 结构体
type CapacityInfo struct {
	// 单位时间段（整点小时纬度，HH:MM-HH:MM） 只能传入单位整点时间段，eg：01:00-02:00
	TimeRange string `json:"time_range,omitempty" xml:"time_range,omitempty"`
	// 时间段产能 （单），表示时间段内可揽可上门单数
	Capacity int64 `json:"capacity,omitempty" xml:"capacity,omitempty"`
}

var poolCapacityInfo = sync.Pool{
	New: func() any {
		return new(CapacityInfo)
	},
}

// GetCapacityInfo() 从对象池中获取CapacityInfo
func GetCapacityInfo() *CapacityInfo {
	return poolCapacityInfo.Get().(*CapacityInfo)
}

// ReleaseCapacityInfo 释放CapacityInfo
func ReleaseCapacityInfo(v *CapacityInfo) {
	v.TimeRange = ""
	v.Capacity = 0
	poolCapacityInfo.Put(v)
}
