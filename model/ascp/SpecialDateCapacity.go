package ascp

import (
	"sync"
)

// SpecialDateCapacity 结构体
type SpecialDateCapacity struct {
	// 日期（年月日），例如20230506
	Date string `json:"date,omitempty" xml:"date,omitempty"`
	// 指定日期下产能（≥0）
	Capacity string `json:"capacity,omitempty" xml:"capacity,omitempty"`
}

var poolSpecialDateCapacity = sync.Pool{
	New: func() any {
		return new(SpecialDateCapacity)
	},
}

// GetSpecialDateCapacity() 从对象池中获取SpecialDateCapacity
func GetSpecialDateCapacity() *SpecialDateCapacity {
	return poolSpecialDateCapacity.Get().(*SpecialDateCapacity)
}

// ReleaseSpecialDateCapacity 释放SpecialDateCapacity
func ReleaseSpecialDateCapacity(v *SpecialDateCapacity) {
	v.Date = ""
	v.Capacity = ""
	poolSpecialDateCapacity.Put(v)
}
