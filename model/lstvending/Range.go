package lstvending

import (
	"sync"
)

// Range 结构体
type Range struct {
	// 结束时间
	End string `json:"end,omitempty" xml:"end,omitempty"`
	// 开始时间
	Begin string `json:"begin,omitempty" xml:"begin,omitempty"`
}

var poolRange = sync.Pool{
	New: func() any {
		return new(Range)
	},
}

// GetRange() 从对象池中获取Range
func GetRange() *Range {
	return poolRange.Get().(*Range)
}

// ReleaseRange 释放Range
func ReleaseRange(v *Range) {
	v.End = ""
	v.Begin = ""
	poolRange.Put(v)
}
