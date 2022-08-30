package refund

// MaxRefundFee 结构体
type MaxRefundFee struct {
	// 可以协商的最大退款金额
	MaxRefundFee int64 `json:"max_refund_fee,omitempty" xml:"max_refund_fee,omitempty"`
}
