package trade

// OrderRefundActionResponse 结构体
type OrderRefundActionResponse struct {
	// 该订单支持的退款退货操作的集合
	SupportRefundActions []Supportrefundactions `json:"support_refund_actions,omitempty" xml:"support_refund_actions>supportrefundactions,omitempty"`
	// 操作用户ID
	OperationUserId string `json:"operation_user_id,omitempty" xml:"operation_user_id,omitempty"`
	// 子订单ID
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
}
