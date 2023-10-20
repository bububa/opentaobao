package alilabs

import (
	"sync"
)

// HashMap 结构体
type HashMap struct {
	// ROM
	Rom string `json:"rom,omitempty" xml:"rom,omitempty"`
	// RAM
	Ram string `json:"ram,omitempty" xml:"ram,omitempty"`
}

var poolHashMap = sync.Pool{
	New: func() any {
		return new(HashMap)
	},
}

// GetHashMap() 从对象池中获取HashMap
func GetHashMap() *HashMap {
	return poolHashMap.Get().(*HashMap)
}

// ReleaseHashMap 释放HashMap
func ReleaseHashMap(v *HashMap) {
	v.Rom = ""
	v.Ram = ""
	poolHashMap.Put(v)
}
