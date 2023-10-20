package omniorder

import (
	"sync"
)

// Store 结构体
type Store struct {
	// 门店库存列表
	StoreInventories []StoreInventory `json:"store_inventories,omitempty" xml:"store_inventories>store_inventory,omitempty"`
	// 库存来源的类型
	WarehouseType string `json:"warehouse_type,omitempty" xml:"warehouse_type,omitempty"`
	// 门店ID(商户中心)或 电商仓ID
	WarehouseId string `json:"warehouse_id,omitempty" xml:"warehouse_id,omitempty"`
}

var poolStore = sync.Pool{
	New: func() any {
		return new(Store)
	},
}

// GetStore() 从对象池中获取Store
func GetStore() *Store {
	return poolStore.Get().(*Store)
}

// ReleaseStore 释放Store
func ReleaseStore(v *Store) {
	v.StoreInventories = v.StoreInventories[:0]
	v.WarehouseType = ""
	v.WarehouseId = ""
	poolStore.Put(v)
}
