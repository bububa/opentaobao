package tmallcar

// ModifyRestPaymentCommand 结构体
type ModifyRestPaymentCommand struct {
	// 订单id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 期望改到的实付金额，单位分
	TargetActualPayFee int64 `json:"target_actual_pay_fee,omitempty" xml:"target_actual_pay_fee,omitempty"`
	// 子支付单id
	SubPayOrderId int64 `json:"sub_pay_order_id,omitempty" xml:"sub_pay_order_id,omitempty"`
}
