package wms

import (
	"sync"
)

// Inventoryitemlist 结构体
type Inventoryitemlist struct {
	// 商品属性列表
	InventoryItem *Inventoryitem `json:"inventory_item,omitempty" xml:"inventory_item,omitempty"`
}

var poolInventoryitemlist = sync.Pool{
	New: func() any {
		return new(Inventoryitemlist)
	},
}

// GetInventoryitemlist() 从对象池中获取Inventoryitemlist
func GetInventoryitemlist() *Inventoryitemlist {
	return poolInventoryitemlist.Get().(*Inventoryitemlist)
}

// ReleaseInventoryitemlist 释放Inventoryitemlist
func ReleaseInventoryitemlist(v *Inventoryitemlist) {
	v.InventoryItem = nil
	poolInventoryitemlist.Put(v)
}
