package ascp

import (
	"sync"
)

// SpecialTimeCapacity 结构体
type SpecialTimeCapacity struct {
	// 开始时间
	StartHour string `json:"start_hour,omitempty" xml:"start_hour,omitempty"`
	// 结束时间
	ToHour string `json:"to_hour,omitempty" xml:"to_hour,omitempty"`
	// 产能
	Capacity string `json:"capacity,omitempty" xml:"capacity,omitempty"`
}

var poolSpecialTimeCapacity = sync.Pool{
	New: func() any {
		return new(SpecialTimeCapacity)
	},
}

// GetSpecialTimeCapacity() 从对象池中获取SpecialTimeCapacity
func GetSpecialTimeCapacity() *SpecialTimeCapacity {
	return poolSpecialTimeCapacity.Get().(*SpecialTimeCapacity)
}

// ReleaseSpecialTimeCapacity 释放SpecialTimeCapacity
func ReleaseSpecialTimeCapacity(v *SpecialTimeCapacity) {
	v.StartHour = ""
	v.ToHour = ""
	v.Capacity = ""
	poolSpecialTimeCapacity.Put(v)
}
