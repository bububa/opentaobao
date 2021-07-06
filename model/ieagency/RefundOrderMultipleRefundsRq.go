package ieagency

// RefundOrderMultipleRefundsRq 结构体
type RefundOrderMultipleRefundsRq struct {
	// 乘机人补退参数
	PassengerMultipleRefundsParams []PassengerMultipleRefundsParam `json:"passenger_multiple_refunds_params,omitempty" xml:"passenger_multiple_refunds_params>passenger_multiple_refunds_param,omitempty"`
	// 代理商ID
	AgentId int64 `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// 退票申请单ID
	RefundOrderId int64 `json:"refund_order_id,omitempty" xml:"refund_order_id,omitempty"`
}
