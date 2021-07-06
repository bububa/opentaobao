package lstwarehouse

// Stocklist 结构体
type Stocklist struct {
	// 商家仓仓库code
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 要设置的库存数量
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 商品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}
