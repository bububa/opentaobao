package icbu

import (
	"sync"
)

// InventoryDetail 结构体
type InventoryDetail struct {
	// 仓库code，默认不填
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 想设置的库存
	CurrentInventory int64 `json:"current_inventory,omitempty" xml:"current_inventory,omitempty"`
	// 原始库存
	SrcInventory int64 `json:"src_inventory,omitempty" xml:"src_inventory,omitempty"`
}

var poolInventoryDetail = sync.Pool{
	New: func() any {
		return new(InventoryDetail)
	},
}

// GetInventoryDetail() 从对象池中获取InventoryDetail
func GetInventoryDetail() *InventoryDetail {
	return poolInventoryDetail.Get().(*InventoryDetail)
}

// ReleaseInventoryDetail 释放InventoryDetail
func ReleaseInventoryDetail(v *InventoryDetail) {
	v.StoreCode = ""
	v.CurrentInventory = 0
	v.SrcInventory = 0
	poolInventoryDetail.Put(v)
}
