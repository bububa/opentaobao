package wms

// Orderitemlist 结构体
type Orderitemlist struct {
	// 订单商品信息
	OrderItem *Orderitem `json:"order_item,omitempty" xml:"order_item,omitempty"`
}
