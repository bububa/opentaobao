package btrip

// BtripFlightPayOrderRs 结构体
type BtripFlightPayOrderRs struct {
	// 最晚支付时间
	LastPayTime string `json:"last_pay_time,omitempty" xml:"last_pay_time,omitempty"`
	// 失败类型
	FailCode string `json:"fail_code,omitempty" xml:"fail_code,omitempty"`
	// 失败原因
	FailReason string `json:"fail_reason,omitempty" xml:"fail_reason,omitempty"`
	// 支付宝流水号（现付支付宝不为空）
	AlipayTradeNo int64 `json:"alipay_trade_no,omitempty" xml:"alipay_trade_no,omitempty"`
	// 实际支付金额
	ActualPayPrice int64 `json:"actual_pay_price,omitempty" xml:"actual_pay_price,omitempty"`
	// 支付状态
	PayStatus int64 `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
}
