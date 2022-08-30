package tmallcar

// CarSubPayOrderDto 结构体
type CarSubPayOrderDto struct {
	// 阶段名称
	StepName string `json:"step_name,omitempty" xml:"step_name,omitempty"`
	// 阶段号
	StepNo int64 `json:"step_no,omitempty" xml:"step_no,omitempty"`
	// 需要付款的金额，单位分
	PayFee int64 `json:"pay_fee,omitempty" xml:"pay_fee,omitempty"`
	// 实际付款金额，单位分
	ActualPayFee int64 `json:"actual_pay_fee,omitempty" xml:"actual_pay_fee,omitempty"`
	// 支付状态
	PayStatus int64 `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
	// 子支付单id
	SubPayOrderId int64 `json:"sub_pay_order_id,omitempty" xml:"sub_pay_order_id,omitempty"`
}
