package security

import (
	"sync"
)

// RpSex 结构体
type RpSex struct {
	// type
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}

var poolRpSex = sync.Pool{
	New: func() any {
		return new(RpSex)
	},
}

// GetRpSex() 从对象池中获取RpSex
func GetRpSex() *RpSex {
	return poolRpSex.Get().(*RpSex)
}

// ReleaseRpSex 释放RpSex
func ReleaseRpSex(v *RpSex) {
	v.Type = 0
	poolRpSex.Put(v)
}
