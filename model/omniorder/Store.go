package omniorder

// Store 结构体
type Store struct {
	// 门店库存列表
	StoreInventories []StoreInventory `json:"store_inventories,omitempty" xml:"store_inventories>store_inventory,omitempty"`
	// 库存来源的类型
	WarehouseType string `json:"warehouse_type,omitempty" xml:"warehouse_type,omitempty"`
	// 门店ID(商户中心)或 电商仓ID
	WarehouseId string `json:"warehouse_id,omitempty" xml:"warehouse_id,omitempty"`
}
