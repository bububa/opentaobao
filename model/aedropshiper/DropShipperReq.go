package aedropshiper

import (
	"sync"
)

// DropShipperReq 结构体
type DropShipperReq struct {
	// Store address
	StoreUrl string `json:"store_url,omitempty" xml:"store_url,omitempty"`
}

var poolDropShipperReq = sync.Pool{
	New: func() any {
		return new(DropShipperReq)
	},
}

// GetDropShipperReq() 从对象池中获取DropShipperReq
func GetDropShipperReq() *DropShipperReq {
	return poolDropShipperReq.Get().(*DropShipperReq)
}

// ReleaseDropShipperReq 释放DropShipperReq
func ReleaseDropShipperReq(v *DropShipperReq) {
	v.StoreUrl = ""
	poolDropShipperReq.Put(v)
}
