package omniorder

// QuantityDetail 结构体
type QuantityDetail struct {
	// 库存类型
	InventoryType string `json:"inventory_type,omitempty" xml:"inventory_type,omitempty"`
	// 库存数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 占用库存
	OccupyQuantity int64 `json:"occupy_quantity,omitempty" xml:"occupy_quantity,omitempty"`
}
