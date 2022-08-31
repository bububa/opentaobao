package idle

// RecycleRefundDetailDto 结构体
type RecycleRefundDetailDto struct {
	// 毫秒，操作超时时间，截止时间
	OpTimeout string `json:"op_timeout,omitempty" xml:"op_timeout,omitempty"`
	// 订单号
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 状态(1, &#34;买家已经申请退款，等待卖家同意&#34;),  (2, &#34;卖家已经同意退货退款，等待买家退货&#34;), --该状态下卖家已经提供退货地址，同意退货退款和发送地址是同时的  (3, &#34;买家已经退货，等待卖家确认收货&#34;),  (4, &#34;退款关闭&#34;),  (5, &#34;退款成功&#34;), --仅退款情况下是从状态1变更到5，同意退款和退款是同时的  (6, &#34;卖家拒绝退款&#34;),  (7, &#34;等待买家确认重新邮寄的货物&#34;),--如补寄等情形，一般不会出现  (8, &#34;等待卖家确认退货地址&#34;)
	RefundStatus string `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	// 退款状态描述
	RefundStatusDesc string `json:"refund_status_desc,omitempty" xml:"refund_status_desc,omitempty"`
	// 退款金额
	RefundFee string `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 退款描述
	RefundDesc string `json:"refund_desc,omitempty" xml:"refund_desc,omitempty"`
	// 申请原因
	ApplyReason string `json:"apply_reason,omitempty" xml:"apply_reason,omitempty"`
	// 拒绝原因
	RefuseReason string `json:"refuse_reason,omitempty" xml:"refuse_reason,omitempty"`
	// 退款完结时间
	RefundEndTime string `json:"refund_end_time,omitempty" xml:"refund_end_time,omitempty"`
	// 退款开始时间
	RefundStartTime string `json:"refund_start_time,omitempty" xml:"refund_start_time,omitempty"`
	// 卖家同意退货说明
	SellerAgreeMsg string `json:"seller_agree_msg,omitempty" xml:"seller_agree_msg,omitempty"`
	// 地址信息，交易状态为2时返回此字段
	SellerAddress *ShippingAddressInfo `json:"seller_address,omitempty" xml:"seller_address,omitempty"`
	// 实际退款金额，不包括追缴单
	ActualRefundFee int64 `json:"actual_refund_fee,omitempty" xml:"actual_refund_fee,omitempty"`
	// 追缴单
	RecoverOrderInfo *AlipayOrderDto `json:"recover_order_info,omitempty" xml:"recover_order_info,omitempty"`
}
