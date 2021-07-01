package wms

// CainiaoStockOutBillInventoryitemlist 结构体
type CainiaoStockOutBillInventoryitemlist struct {
	// 商品属性
	InventoryItem *CainiaoStockOutBillInventoryitem `json:"inventory_item,omitempty" xml:"inventory_item,omitempty"`
}
