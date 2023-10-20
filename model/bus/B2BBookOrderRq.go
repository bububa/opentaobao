package bus

// B2bbookOrderRq 结构体
type B2bbookOrderRq struct {
	// 阿里订单号
	AliOrderId string `json:"ali_order_id,omitempty" xml:"ali_order_id,omitempty"`
}
