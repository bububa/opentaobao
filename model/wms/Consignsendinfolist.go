package wms

import (
	"sync"
)

// Consignsendinfolist 结构体
type Consignsendinfolist struct {
	// 物流订单发货信息
	ConsignSendInfo *Consignsendinfo `json:"consign_send_info,omitempty" xml:"consign_send_info,omitempty"`
}

var poolConsignsendinfolist = sync.Pool{
	New: func() any {
		return new(Consignsendinfolist)
	},
}

// GetConsignsendinfolist() 从对象池中获取Consignsendinfolist
func GetConsignsendinfolist() *Consignsendinfolist {
	return poolConsignsendinfolist.Get().(*Consignsendinfolist)
}

// ReleaseConsignsendinfolist 释放Consignsendinfolist
func ReleaseConsignsendinfolist(v *Consignsendinfolist) {
	v.ConsignSendInfo = nil
	poolConsignsendinfolist.Put(v)
}
