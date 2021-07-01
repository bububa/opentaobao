package wms

// Orderitemlistwlbwmsstockoutordernotify 结构体
type Orderitemlistwlbwmsstockoutordernotify struct {
	// 订单商品信息
	OrderItem *Orderitemwlbwmsstockoutordernotify `json:"order_item,omitempty" xml:"order_item,omitempty"`
}
