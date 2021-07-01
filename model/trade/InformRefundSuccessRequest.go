package trade

// InformRefundSuccessRequest 结构体
type InformRefundSuccessRequest struct {
	// 门店编码
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 退货成功履约子单
	RefundFulfillSubOrders []RefundSuccessSubOrderRequest `json:"refund_fulfill_sub_orders,omitempty" xml:"refund_fulfill_sub_orders>refund_success_sub_order_request,omitempty"`
}
