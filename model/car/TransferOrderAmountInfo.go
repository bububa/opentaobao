package car

// TransferOrderAmountInfo 结构体
type TransferOrderAmountInfo struct {
	// 总价(单位 元)
	TotalFee string `json:"total_fee,omitempty" xml:"total_fee,omitempty"`
	// 实际线上支付金额(单位 元)
	UserRealPay string `json:"user_real_pay,omitempty" xml:"user_real_pay,omitempty"`
	// 实际付给商家的钱(单位 元)
	Payment string `json:"payment,omitempty" xml:"payment,omitempty"`
	// 退款金额(单位 元)
	RefundFee string `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
}
