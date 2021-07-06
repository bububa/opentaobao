package xhotelonlineorder

// TopPostTradeRefundDo 结构体
type TopPostTradeRefundDo struct {
	// 退款成功时间
	RefundSucTime string `json:"refund_suc_time,omitempty" xml:"refund_suc_time,omitempty"`
	// 创建退款申请时的额外描述
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 引起退款申请的类型描述
	RefundReason string `json:"refund_reason,omitempty" xml:"refund_reason,omitempty"`
	// 退款发起时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 申请的退款金额（单位：分）
	RefundAmount int64 `json:"refund_amount,omitempty" xml:"refund_amount,omitempty"`
	// 售后退款类型,1:买家取消订单，创建退款申请;2:卖家发起协商取消;3:卖家用户协商一致后，小二执行退款操作;4:卖家发起协商取消退款,系统收到买家确认短信后进行退款
	PostTradeRefundType int64 `json:"post_trade_refund_type,omitempty" xml:"post_trade_refund_type,omitempty"`
	// 退款状态,1:退款申请已创建;2:退款申请已拒绝;3:退款处理中;4:退款已成功;5:退款失败;6: 退款申请已取消
	PostTradeRefundStatus int64 `json:"post_trade_refund_status,omitempty" xml:"post_trade_refund_status,omitempty"`
	// tid
	Tid int64 `json:"tid,omitempty" xml:"tid,omitempty"`
}
