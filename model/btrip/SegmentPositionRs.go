package btrip

import (
	"sync"
)

// SegmentPositionRs 结构体
type SegmentPositionRs struct {
}

var poolSegmentPositionRs = sync.Pool{
	New: func() any {
		return new(SegmentPositionRs)
	},
}

// GetSegmentPositionRs() 从对象池中获取SegmentPositionRs
func GetSegmentPositionRs() *SegmentPositionRs {
	return poolSegmentPositionRs.Get().(*SegmentPositionRs)
}

// ReleaseSegmentPositionRs 释放SegmentPositionRs
func ReleaseSegmentPositionRs(v *SegmentPositionRs) {
	poolSegmentPositionRs.Put(v)
}
