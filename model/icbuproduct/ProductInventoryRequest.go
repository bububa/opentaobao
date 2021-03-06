package icbuproduct

// ProductInventoryRequest 结构体
type ProductInventoryRequest struct {
	// 待更新的库存列表
	InventoryList []InventoryDto `json:"inventory_list,omitempty" xml:"inventory_list>inventory_dto,omitempty"`
	// 商品id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
}
