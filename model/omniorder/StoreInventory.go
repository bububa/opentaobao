package omniorder

// StoreInventory 结构体
type StoreInventory struct {
	// 库存量详情列表
	QuantityDetails []QuantityDetail `json:"quantity_details,omitempty" xml:"quantity_details>quantity_detail,omitempty"`
	// 单据流水号，用于幂等操作
	BillNum string `json:"bill_num,omitempty" xml:"bill_num,omitempty"`
	// ISV系统中商品编码
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 库存类型
	InventoryType string `json:"inventory_type,omitempty" xml:"inventory_type,omitempty"`
	// 淘宝前端商品id
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 商品的SKU编码
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 单据类型
	BillType string `json:"bill_type,omitempty" xml:"bill_type,omitempty"`
	// 淘宝后端商品id
	ScItemId int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// 对应类型的库存数量（正数）
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 调整后库存数量
	FinalQuantity int64 `json:"final_quantity,omitempty" xml:"final_quantity,omitempty"`
}
