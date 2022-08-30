package examination

// OrderRefundVO 结构体
type OrderRefundVO struct {
	// 子订单信息
	SubOrderList []SubOrderRequest `json:"sub_order_list,omitempty" xml:"sub_order_list>sub_order_request,omitempty"`
	// 主订单
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 退款id
	RefundOrderId string `json:"refund_order_id,omitempty" xml:"refund_order_id,omitempty"`
}
