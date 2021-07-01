package wms

// CainiaoReturnBillOrderitemlist 结构体
type CainiaoReturnBillOrderitemlist struct {
	// 订单商品信息
	OrderItem *CainiaoReturnBillOrderitem `json:"order_item,omitempty" xml:"order_item,omitempty"`
}
