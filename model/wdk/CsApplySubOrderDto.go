package wdk

// CsApplySubOrderDto 结构体
type CsApplySubOrderDto struct {
	// 渠道子订单号，淘鲜达渠道为TP子单号
	OutSubOrderId string `json:"out_sub_order_id,omitempty" xml:"out_sub_order_id,omitempty"`
	// 申请子单退款金额
	RefundFee *BigDecimal `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 申请子单退货数量
	RefundAmount *BigDecimal `json:"refund_amount,omitempty" xml:"refund_amount,omitempty"`
}
