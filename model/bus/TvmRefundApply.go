package bus

// TvmRefundApply 结构体
type TvmRefundApply struct {
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 成功时间
	GmtRefundSuccTime string `json:"gmt_refund_succ_time,omitempty" xml:"gmt_refund_succ_time,omitempty"`
	// 退款批次号
	OutTradeNo string `json:"out_trade_no,omitempty" xml:"out_trade_no,omitempty"`
	// 申请单id
	ApplyId int64 `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 退款金额（分）
	RefundAmount int64 `json:"refund_amount,omitempty" xml:"refund_amount,omitempty"`
	// 退款状态 10(处理中) 20(已拒绝) 30(已同意) 40(已退款) 50(已受理)
	RefundStatus int64 `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	// 淘宝退款申请单id
	RpApplyId int64 `json:"rp_apply_id,omitempty" xml:"rp_apply_id,omitempty"`
}
