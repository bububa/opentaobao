package store

import (
	"sync"
)

// Models 结构体
type Models struct {
	// 是否为空
	Empty bool `json:"empty,omitempty" xml:"empty,omitempty"`
}

var poolModels = sync.Pool{
	New: func() any {
		return new(Models)
	},
}

// GetModels() 从对象池中获取Models
func GetModels() *Models {
	return poolModels.Get().(*Models)
}

// ReleaseModels 释放Models
func ReleaseModels(v *Models) {
	v.Empty = false
	poolModels.Put(v)
}
