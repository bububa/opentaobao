package wms

// CainiaoStockInBillInventoryitemlist 结构体
type CainiaoStockInBillInventoryitemlist struct {
	// 仓库收货商品信息
	InventoryItem *CainiaoStockInBillInventoryitem `json:"inventory_item,omitempty" xml:"inventory_item,omitempty"`
}
