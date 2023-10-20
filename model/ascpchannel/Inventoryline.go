package ascpchannel

import (
	"sync"
)

// Inventoryline 结构体
type Inventoryline struct {
	// 库存数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}

var poolInventoryline = sync.Pool{
	New: func() any {
		return new(Inventoryline)
	},
}

// GetInventoryline() 从对象池中获取Inventoryline
func GetInventoryline() *Inventoryline {
	return poolInventoryline.Get().(*Inventoryline)
}

// ReleaseInventoryline 释放Inventoryline
func ReleaseInventoryline(v *Inventoryline) {
	v.Quantity = 0
	poolInventoryline.Put(v)
}
