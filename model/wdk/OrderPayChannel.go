package wdk

// OrderPayChannel 结构体
type OrderPayChannel struct {
	// 支付方式
	PayType string `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	// 支付渠道名称
	PayChannel string `json:"pay_channel,omitempty" xml:"pay_channel,omitempty"`
	// 支付金额
	PayFee int64 `json:"pay_fee,omitempty" xml:"pay_fee,omitempty"`
}
