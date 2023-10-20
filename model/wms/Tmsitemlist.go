package wms

import (
	"sync"
)

// Tmsitemlist 结构体
type Tmsitemlist struct {
	// 包裹里面商品
	TmsItem *Tmsitem `json:"tms_item,omitempty" xml:"tms_item,omitempty"`
}

var poolTmsitemlist = sync.Pool{
	New: func() any {
		return new(Tmsitemlist)
	},
}

// GetTmsitemlist() 从对象池中获取Tmsitemlist
func GetTmsitemlist() *Tmsitemlist {
	return poolTmsitemlist.Get().(*Tmsitemlist)
}

// ReleaseTmsitemlist 释放Tmsitemlist
func ReleaseTmsitemlist(v *Tmsitemlist) {
	v.TmsItem = nil
	poolTmsitemlist.Put(v)
}
