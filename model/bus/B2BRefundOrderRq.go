package bus

// B2brefundOrderRq 结构体
type B2brefundOrderRq struct {
	// 飞猪子订单号
	SubOrderIds []int64 `json:"sub_order_ids,omitempty" xml:"sub_order_ids>int64,omitempty"`
	// 退票原因
	RefundReason string `json:"refund_reason,omitempty" xml:"refund_reason,omitempty"`
	// 飞猪订单号
	AliTripOrderId string `json:"ali_trip_order_id,omitempty" xml:"ali_trip_order_id,omitempty"`
	// 卖家淘宝ID
	SellerAgentId int64 `json:"seller_agent_id,omitempty" xml:"seller_agent_id,omitempty"`
}
