package btrip

import (
	"sync"
)

// SegmentShowInfoRs 结构体
type SegmentShowInfoRs struct {
}

var poolSegmentShowInfoRs = sync.Pool{
	New: func() any {
		return new(SegmentShowInfoRs)
	},
}

// GetSegmentShowInfoRs() 从对象池中获取SegmentShowInfoRs
func GetSegmentShowInfoRs() *SegmentShowInfoRs {
	return poolSegmentShowInfoRs.Get().(*SegmentShowInfoRs)
}

// ReleaseSegmentShowInfoRs 释放SegmentShowInfoRs
func ReleaseSegmentShowInfoRs(v *SegmentShowInfoRs) {
	poolSegmentShowInfoRs.Put(v)
}
