package wdk

// RefundPayChannel 结构体
type RefundPayChannel struct {
	// 支付渠道类型 1.默认 10-支付宝 20-微信支付 30-积分支付 40-储值卡支付 50-银行卡支付。有支付渠道的情况下，必填。 必须是翱象支持的支付渠道，否则报错。
	PayChannelType int64 `json:"pay_channel_type,omitempty" xml:"pay_channel_type,omitempty"`
	// 当前支付渠道的退款金额，单位分。有支付渠道的情况下，必填。
	RefundFee int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
}
