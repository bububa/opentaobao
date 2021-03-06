package tmallnr

// StoreDto 结构体
type StoreDto struct {
	// 门店商品，最大列表长度：20
	StoreInventories []StoreInvetoryDto `json:"store_inventories,omitempty" xml:"store_inventories>store_invetory_dto,omitempty"`
	// 库存来源的类型；STORE表示门店
	WarehouseType string `json:"warehouse_type,omitempty" xml:"warehouse_type,omitempty"`
	// 门店对应的storeID值
	WarehouseId string `json:"warehouse_id,omitempty" xml:"warehouse_id,omitempty"`
}
