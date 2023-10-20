package wms

import (
	"sync"
)

// Sninfolist 结构体
type Sninfolist struct {
	// SN信息
	SnInfo *Sninfo `json:"sn_info,omitempty" xml:"sn_info,omitempty"`
}

var poolSninfolist = sync.Pool{
	New: func() any {
		return new(Sninfolist)
	},
}

// GetSninfolist() 从对象池中获取Sninfolist
func GetSninfolist() *Sninfolist {
	return poolSninfolist.Get().(*Sninfolist)
}

// ReleaseSninfolist 释放Sninfolist
func ReleaseSninfolist(v *Sninfolist) {
	v.SnInfo = nil
	poolSninfolist.Put(v)
}
