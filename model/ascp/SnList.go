package ascp

import (
	"sync"
)

// SnList 结构体
type SnList struct {
	// sn
	Sn []string `json:"sn,omitempty" xml:"sn>string,omitempty"`
}

var poolSnList = sync.Pool{
	New: func() any {
		return new(SnList)
	},
}

// GetSnList() 从对象池中获取SnList
func GetSnList() *SnList {
	return poolSnList.Get().(*SnList)
}

// ReleaseSnList 释放SnList
func ReleaseSnList(v *SnList) {
	v.Sn = v.Sn[:0]
	poolSnList.Put(v)
}
