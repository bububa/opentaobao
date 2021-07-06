package ieagency

// PassengerMultipleRefundsParam 结构体
type PassengerMultipleRefundsParam struct {
	// 乘机人姓名
	PassengerName string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// 补退金额(单位:分)
	RefundMoney int64 `json:"refund_money,omitempty" xml:"refund_money,omitempty"`
}
