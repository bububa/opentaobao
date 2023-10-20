package wms

import (
	"sync"
)

// WmsInventoryQueryItemlist 结构体
type WmsInventoryQueryItemlist struct {
	// 商品详情
	Item *WmsInventoryQueryItem `json:"item,omitempty" xml:"item,omitempty"`
}

var poolWmsInventoryQueryItemlist = sync.Pool{
	New: func() any {
		return new(WmsInventoryQueryItemlist)
	},
}

// GetWmsInventoryQueryItemlist() 从对象池中获取WmsInventoryQueryItemlist
func GetWmsInventoryQueryItemlist() *WmsInventoryQueryItemlist {
	return poolWmsInventoryQueryItemlist.Get().(*WmsInventoryQueryItemlist)
}

// ReleaseWmsInventoryQueryItemlist 释放WmsInventoryQueryItemlist
func ReleaseWmsInventoryQueryItemlist(v *WmsInventoryQueryItemlist) {
	v.Item = nil
	poolWmsInventoryQueryItemlist.Put(v)
}
