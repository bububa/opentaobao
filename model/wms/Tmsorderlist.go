package wms

import (
	"sync"
)

// Tmsorderlist 结构体
type Tmsorderlist struct {
	// 运单信息列表
	TmsOrder *Tmsorder `json:"tms_order,omitempty" xml:"tms_order,omitempty"`
}

var poolTmsorderlist = sync.Pool{
	New: func() any {
		return new(Tmsorderlist)
	},
}

// GetTmsorderlist() 从对象池中获取Tmsorderlist
func GetTmsorderlist() *Tmsorderlist {
	return poolTmsorderlist.Get().(*Tmsorderlist)
}

// ReleaseTmsorderlist 释放Tmsorderlist
func ReleaseTmsorderlist(v *Tmsorderlist) {
	v.TmsOrder = nil
	poolTmsorderlist.Put(v)
}
