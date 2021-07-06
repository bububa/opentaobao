package trade

// RefundReq 结构体
type RefundReq struct {
	// 代理商订单号
	AgentOrderId string `json:"agent_order_id,omitempty" xml:"agent_order_id,omitempty"`
	// 退票失败码
	FailCode string `json:"fail_code,omitempty" xml:"fail_code,omitempty"`
	// 退票失败原因
	FailReason string `json:"fail_reason,omitempty" xml:"fail_reason,omitempty"`
	// 代理商id
	AgentId int64 `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// 退票申请单号
	ApplyId int64 `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 手续费(单位分)
	ChargeFee int64 `json:"charge_fee,omitempty" xml:"charge_fee,omitempty"`
	// 应退额(单位分)
	RefundFee int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 平台订单号
	TpOrderId int64 `json:"tp_order_id,omitempty" xml:"tp_order_id,omitempty"`
	// 退票成功标志
	Status bool `json:"status,omitempty" xml:"status,omitempty"`
}
