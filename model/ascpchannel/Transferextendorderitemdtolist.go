package ascpchannel

import (
	"sync"
)

// Transferextendorderitemdtolist 结构体
type Transferextendorderitemdtolist struct {
	// 批次号
	BatchCode string `json:"batch_code,omitempty" xml:"batch_code,omitempty"`
	// 库存类型
	InventoryType int64 `json:"inventory_type,omitempty" xml:"inventory_type,omitempty"`
	// 数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}

var poolTransferextendorderitemdtolist = sync.Pool{
	New: func() any {
		return new(Transferextendorderitemdtolist)
	},
}

// GetTransferextendorderitemdtolist() 从对象池中获取Transferextendorderitemdtolist
func GetTransferextendorderitemdtolist() *Transferextendorderitemdtolist {
	return poolTransferextendorderitemdtolist.Get().(*Transferextendorderitemdtolist)
}

// ReleaseTransferextendorderitemdtolist 释放Transferextendorderitemdtolist
func ReleaseTransferextendorderitemdtolist(v *Transferextendorderitemdtolist) {
	v.BatchCode = ""
	v.InventoryType = 0
	v.Quantity = 0
	poolTransferextendorderitemdtolist.Put(v)
}
