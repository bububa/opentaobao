package ieagency

// RefundOrderSimpleVo 结构体
type RefundOrderSimpleVo struct {
	// 代理商ID
	AgentId int64 `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// 新老模型（V1:老模型，V2：新模型）
	ModelVersion string `json:"model_version,omitempty" xml:"model_version,omitempty"`
	// 订单ID
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 10: "已提交待处理",WAI20:"待人工处理"
	RefundBizStatus int64 `json:"refund_biz_status,omitempty" xml:"refund_biz_status,omitempty"`
	// 退票单ID
	RefundOrderId int64 `json:"refund_order_id,omitempty" xml:"refund_order_id,omitempty"`
	// 申请单状态(WAIT(1,"待处理"), AGREED(2, "已同意"),REFUSE(3, "已拒绝"),PROCESS(6, "已受理"), SUCCESS(7, "已退款"))
	RefundOrderStatus int64 `json:"refund_order_status,omitempty" xml:"refund_order_status,omitempty"`
	// 申请单支付状态(   INIT(1, "初始化"),     REFUND_FAIL(2, "退款失败"),     REFUND_SUCCESS(3, "退款成功"))
	RefundPayStatus int64 `json:"refund_pay_status,omitempty" xml:"refund_pay_status,omitempty"`
}
