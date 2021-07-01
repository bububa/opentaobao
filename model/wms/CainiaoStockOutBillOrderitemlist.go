package wms

// CainiaoStockOutBillOrderitemlist 结构体
type CainiaoStockOutBillOrderitemlist struct {
	// 订单商品信息
	OrderItem *CainiaoStockOutBillOrderitem `json:"order_item,omitempty" xml:"order_item,omitempty"`
}
