package wms

// CainiaoReturnBillInventoryitemlist 结构体
type CainiaoReturnBillInventoryitemlist struct {
	// 订单详情
	InventoryItem *CainiaoReturnBillInventoryitem `json:"inventory_item,omitempty" xml:"inventory_item,omitempty"`
}
