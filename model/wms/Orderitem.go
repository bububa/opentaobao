package wms

// Orderitem 结构体
type Orderitem struct {
	// 订单明细行编码
	OrderItemId string `json:"order_item_id,omitempty" xml:"order_item_id,omitempty"`
	// 交易单号
	OrderSourceCode string `json:"order_source_code,omitempty" xml:"order_source_code,omitempty"`
	// 商品ID
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 商家编码
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// 商品属性列表
	InventoryItemList []Inventoryitemlist `json:"inventory_item_list,omitempty" xml:"inventory_item_list>inventoryitemlist,omitempty"`
}
