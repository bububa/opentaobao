package logistic

import (
	"sync"
)

// Range 结构体
type Range struct {
	// 经度
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 纬度
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
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
	v.Longitude = ""
	v.Latitude = ""
	poolRange.Put(v)
}
