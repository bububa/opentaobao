package promotion

import (
	"sync"
)

// Extmap 结构体
type Extmap struct {
	// 扩展字段
	Key string `json:"key,omitempty" xml:"key,omitempty"`
	// 扩展字段
	Keyvalue bool `json:"keyvalue,omitempty" xml:"keyvalue,omitempty"`
	// empty
	Empty bool `json:"empty,omitempty" xml:"empty,omitempty"`
}

var poolExtmap = sync.Pool{
	New: func() any {
		return new(Extmap)
	},
}

// GetExtmap() 从对象池中获取Extmap
func GetExtmap() *Extmap {
	return poolExtmap.Get().(*Extmap)
}

// ReleaseExtmap 释放Extmap
func ReleaseExtmap(v *Extmap) {
	v.Key = ""
	v.Keyvalue = false
	v.Empty = false
	poolExtmap.Put(v)
}
