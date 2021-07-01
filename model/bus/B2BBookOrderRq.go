package bus

// B2BBookOrderRq 结构体
type B2BBookOrderRq struct {
	// 阿里订单号
	AliOrderId string `json:"ali_order_id,omitempty" xml:"ali_order_id,omitempty"`
}
