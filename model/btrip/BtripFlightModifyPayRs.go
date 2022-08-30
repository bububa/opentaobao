package btrip

// BtripFlightModifyPayRs 结构体
type BtripFlightModifyPayRs struct {
	// 支付状态
	PayStatus string `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
	// 支付时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 支付宝交易流水号
	TradeNo string `json:"trade_no,omitempty" xml:"trade_no,omitempty"`
	// 支付价格
	PayPrice int64 `json:"pay_price,omitempty" xml:"pay_price,omitempty"`
	// 能否重试
	CanRetry bool `json:"can_retry,omitempty" xml:"can_retry,omitempty"`
}
