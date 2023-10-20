package store

import (
	"sync"
)

// Other 结构体
type Other struct {
	// 是否为空
	Empty bool `json:"empty,omitempty" xml:"empty,omitempty"`
}

var poolOther = sync.Pool{
	New: func() any {
		return new(Other)
	},
}

// GetOther() 从对象池中获取Other
func GetOther() *Other {
	return poolOther.Get().(*Other)
}

// ReleaseOther 释放Other
func ReleaseOther(v *Other) {
	v.Empty = false
	poolOther.Put(v)
}
