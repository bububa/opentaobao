package wenyuvideo

import (
	"sync"
)

// Segments 结构体
type Segments struct {
	// 开始时间点
	From int64 `json:"from,omitempty" xml:"from,omitempty"`
	// 结束时间点
	To int64 `json:"to,omitempty" xml:"to,omitempty"`
}

var poolSegments = sync.Pool{
	New: func() any {
		return new(Segments)
	},
}

// GetSegments() 从对象池中获取Segments
func GetSegments() *Segments {
	return poolSegments.Get().(*Segments)
}

// ReleaseSegments 释放Segments
func ReleaseSegments(v *Segments) {
	v.From = 0
	v.To = 0
	poolSegments.Put(v)
}
