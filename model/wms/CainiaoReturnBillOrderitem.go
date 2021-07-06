package wms

// CainiaoReturnBillOrderitem 结构体
type CainiaoReturnBillOrderitem struct {
	// 商品信息
	InventoryItemList []CainiaoReturnBillInventoryitemlist `json:"inventory_item_list,omitempty" xml:"inventory_item_list>cainiao_return_bill_inventoryitemlist,omitempty"`
	// 商品ID
	OrderItemId string `json:"order_item_id,omitempty" xml:"order_item_id,omitempty"`
	// 商品ID
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 商家编码
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
}
