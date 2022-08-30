package alsc

// RefundOpenInfo 结构体
type RefundOpenInfo struct {
	// 退款单业务上下文
	BizContext string `json:"biz_context,omitempty" xml:"biz_context,omitempty"`
	// 退款单扩展信息
	ExtInfo string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
	// 退款单id
	RefundOrderId string `json:"refund_order_id,omitempty" xml:"refund_order_id,omitempty"`
	// 退款状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 退款单资金明细
	FundOpenInfo *FundOpenInfo `json:"fund_open_info,omitempty" xml:"fund_open_info,omitempty"`
	// 退款金额
	RefundAmount int64 `json:"refund_amount,omitempty" xml:"refund_amount,omitempty"`
	// 退款实付
	RefundRealAmount int64 `json:"refund_real_amount,omitempty" xml:"refund_real_amount,omitempty"`
}
