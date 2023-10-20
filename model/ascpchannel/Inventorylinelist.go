package ascpchannel

import (
	"sync"
)

// Inventorylinelist 结构体
type Inventorylinelist struct {
	// 在库库存操作行对象
	InventoryLine *Inventoryline `json:"inventory_line,omitempty" xml:"inventory_line,omitempty"`
}

var poolInventorylinelist = sync.Pool{
	New: func() any {
		return new(Inventorylinelist)
	},
}

// GetInventorylinelist() 从对象池中获取Inventorylinelist
func GetInventorylinelist() *Inventorylinelist {
	return poolInventorylinelist.Get().(*Inventorylinelist)
}

// ReleaseInventorylinelist 释放Inventorylinelist
func ReleaseInventorylinelist(v *Inventorylinelist) {
	v.InventoryLine = nil
	poolInventorylinelist.Put(v)
}
