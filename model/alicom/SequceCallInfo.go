package alicom

import (
	"sync"
)

// SequceCallInfo 结构体
type SequceCallInfo struct {
	// ff
	SequenceCallEvents []SequenceCallEvent `json:"sequence_call_events,omitempty" xml:"sequence_call_events>sequence_call_event,omitempty"`
	// 顺振call的次数
	SequenceCallCount int64 `json:"sequence_call_count,omitempty" xml:"sequence_call_count,omitempty"`
}

var poolSequceCallInfo = sync.Pool{
	New: func() any {
		return new(SequceCallInfo)
	},
}

// GetSequceCallInfo() 从对象池中获取SequceCallInfo
func GetSequceCallInfo() *SequceCallInfo {
	return poolSequceCallInfo.Get().(*SequceCallInfo)
}

// ReleaseSequceCallInfo 释放SequceCallInfo
func ReleaseSequceCallInfo(v *SequceCallInfo) {
	v.SequenceCallEvents = v.SequenceCallEvents[:0]
	v.SequenceCallCount = 0
	poolSequceCallInfo.Put(v)
}
