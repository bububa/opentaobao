package axindata

// CancelPolicyDetailDto 结构体
type CancelPolicyDetailDto struct {
	// 扣除值
	DeductFeeRate int64 `json:"deduct_fee_rate,omitempty" xml:"deduct_fee_rate,omitempty"`
	// 提前小时数
	AheadCancelHours int64 `json:"ahead_cancel_hours,omitempty" xml:"ahead_cancel_hours,omitempty"`
}
