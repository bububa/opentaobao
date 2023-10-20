package alsc

import (
	"sync"
)

// Extinfo 结构体
type Extinfo struct {
	// Map&lt;String, String&gt;
	Empty bool `json:"empty,omitempty" xml:"empty,omitempty"`
}

var poolExtinfo = sync.Pool{
	New: func() any {
		return new(Extinfo)
	},
}

// GetExtinfo() 从对象池中获取Extinfo
func GetExtinfo() *Extinfo {
	return poolExtinfo.Get().(*Extinfo)
}

// ReleaseExtinfo 释放Extinfo
func ReleaseExtinfo(v *Extinfo) {
	v.Empty = false
	poolExtinfo.Put(v)
}
