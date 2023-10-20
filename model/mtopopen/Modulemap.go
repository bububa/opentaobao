package mtopopen

import (
	"sync"
)

// Modulemap 结构体
type Modulemap struct {
	// false
	Empty bool `json:"empty,omitempty" xml:"empty,omitempty"`
}

var poolModulemap = sync.Pool{
	New: func() any {
		return new(Modulemap)
	},
}

// GetModulemap() 从对象池中获取Modulemap
func GetModulemap() *Modulemap {
	return poolModulemap.Get().(*Modulemap)
}

// ReleaseModulemap 释放Modulemap
func ReleaseModulemap(v *Modulemap) {
	v.Empty = false
	poolModulemap.Put(v)
}
