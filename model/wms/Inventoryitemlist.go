package wms

// Inventoryitemlist 结构体
type Inventoryitemlist struct {
	// 商品属性列表
	InventoryItem *Inventoryitem `json:"inventory_item,omitempty" xml:"inventory_item,omitempty"`
}
