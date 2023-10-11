package ieagency

// RefundOrderVo 结构体
type RefundOrderVo struct {
	// 乘机人费用列表
	RefundPassengerFeeVos []RefundPassengerFeeVo `json:"refund_passenger_fee_vos,omitempty" xml:"refund_passenger_fee_vos>refund_passenger_fee_vo,omitempty"`
	// 退票乘机人列表
	RefundPassengerVos []RefundPassengerVo `json:"refund_passenger_vos,omitempty" xml:"refund_passenger_vos>refund_passenger_vo,omitempty"`
	// 代理商ID
	AgentId int64 `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// 正向订单ID
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 业务状态(INIT(10, &#34;已提交待处理&#34;),     WAIT_ADMIN_PROCESS(20, &#34;待人工处理&#34;)),配合主状态使用
	RefundBizStatus int64 `json:"refund_biz_status,omitempty" xml:"refund_biz_status,omitempty"`
	// 退商品信息
	RefundItemVo *RefundItemVo `json:"refund_item_vo,omitempty" xml:"refund_item_vo,omitempty"`
	// 申请单详情
	RefundOrderDetailVo *RefundOrderDetailVo `json:"refund_order_detail_vo,omitempty" xml:"refund_order_detail_vo,omitempty"`
	// 退票申请单ID
	RefundOrderId int64 `json:"refund_order_id,omitempty" xml:"refund_order_id,omitempty"`
	// 申请单状态(WAIT(1,&#34;待处理&#34;), AGREED(2, &#34;已同意&#34;),REFUSE(3, &#34;已拒绝&#34;),PROCESS(6, &#34;已受理&#34;), SUCCESS(7, &#34;已退款&#34;)
	RefundOrderStatus int64 `json:"refund_order_status,omitempty" xml:"refund_order_status,omitempty"`
	// 支付状态(INIT(1, &#34;初始化&#34;),     REFUND_FAIL(2, &#34;退款失败&#34;),     REFUND_SUCCESS(3, &#34;退款成功&#34;)
	RefundPayStatus int64 `json:"refund_pay_status,omitempty" xml:"refund_pay_status,omitempty"`
	// 原因
	RefundReasonDo *IeRefundReasonDo `json:"refund_reason_do,omitempty" xml:"refund_reason_do,omitempty"`
	// 总退给买家金额(单位:分)
	TotalRefundToBuyerMoney int64 `json:"total_refund_to_buyer_money,omitempty" xml:"total_refund_to_buyer_money,omitempty"`
	// 总活动收回金额(单位:分)
	TotalTakeBackActivityMoney int64 `json:"total_take_back_activity_money,omitempty" xml:"total_take_back_activity_money,omitempty"`
	// 改签信息
	ChangeSimpleVo *ChangeSimpleVo `json:"change_simple_vo,omitempty" xml:"change_simple_vo,omitempty"`
	// 催退
	UrgentRefundVo *UrgentRefundVo `json:"urgent_refund_vo,omitempty" xml:"urgent_refund_vo,omitempty"`
}
