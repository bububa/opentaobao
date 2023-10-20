package ascp

import (
	"sync"
)

// SpecifyCapacityInfo 结构体
type SpecifyCapacityInfo struct {
	// 时间段产能
	SpecifyDateCapacity []CapacityInfo `json:"specify_date_capacity,omitempty" xml:"specify_date_capacity>capacity_info,omitempty"`
	// 指定日期，YYYY-MM-DD
	SpecifyDate string `json:"specify_date,omitempty" xml:"specify_date,omitempty"`
}

var poolSpecifyCapacityInfo = sync.Pool{
	New: func() any {
		return new(SpecifyCapacityInfo)
	},
}

// GetSpecifyCapacityInfo() 从对象池中获取SpecifyCapacityInfo
func GetSpecifyCapacityInfo() *SpecifyCapacityInfo {
	return poolSpecifyCapacityInfo.Get().(*SpecifyCapacityInfo)
}

// ReleaseSpecifyCapacityInfo 释放SpecifyCapacityInfo
func ReleaseSpecifyCapacityInfo(v *SpecifyCapacityInfo) {
	v.SpecifyDateCapacity = v.SpecifyDateCapacity[:0]
	v.SpecifyDate = ""
	poolSpecifyCapacityInfo.Put(v)
}
