package wms

// CainiaoStockInBillOrderitemlist 结构体
type CainiaoStockInBillOrderitemlist struct {
	// 入库单信息
	OrderItem *CainiaoStockInBillOrderitem `json:"order_item,omitempty" xml:"order_item,omitempty"`
}
