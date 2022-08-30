package btrip

// BtripFlightRefundApplyRs 结构体
type BtripFlightRefundApplyRs struct {
	// 分销外部订单号
	DisOrderId string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
	// 分销外部退票单号
	DisSubOrderId string `json:"dis_sub_order_id,omitempty" xml:"dis_sub_order_id,omitempty"`
	// 退票手续费
	RefundFee int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 退票应退
	RefundMoney int64 `json:"refund_money,omitempty" xml:"refund_money,omitempty"`
	// 商旅退票申请单号
	RefundApplyId int64 `json:"refund_apply_id,omitempty" xml:"refund_apply_id,omitempty"`
}
