package axintrade

// ConditionRefundPolicyDto 结构体
type ConditionRefundPolicyDto struct {
	// 条件退类型
	ConditionRefundType int64 `json:"condition_refund_type,omitempty" xml:"condition_refund_type,omitempty"`
	// 允许退款的最后时间点天数
	Days int64 `json:"days,omitempty" xml:"days,omitempty"`
	// 允许退款的最后的时间点小时
	Hours int64 `json:"hours,omitempty" xml:"hours,omitempty"`
	// 允许退款的最后的时间点分钟
	Minutes int64 `json:"minutes,omitempty" xml:"minutes,omitempty"`
	// 退款手续费金额，单位分
	ChargeAmount int64 `json:"charge_amount,omitempty" xml:"charge_amount,omitempty"`
	// 退款手续费比例
	Percent int64 `json:"percent,omitempty" xml:"percent,omitempty"`
}
