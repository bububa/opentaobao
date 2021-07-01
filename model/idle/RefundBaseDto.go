package idle

// RefundBaseDto 结构体
type RefundBaseDto struct {
	// 订单号
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 退款状态
	RefundStatus int64 `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	// 退款状态描述
	RefundStatusDesc string `json:"refund_status_desc,omitempty" xml:"refund_status_desc,omitempty"`
	// 退款金额/分
	RefundFee int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
}
