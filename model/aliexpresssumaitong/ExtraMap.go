package aliexpresssumaitong

// ExtraMap 结构体
type ExtraMap struct {
	// 支付收单号
	PaymentId string `json:"payment_id,omitempty" xml:"payment_id,omitempty"`
	// 订单状态
	OrderStatus string `json:"order_status,omitempty" xml:"order_status,omitempty"`
}
