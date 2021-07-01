package drug

// FundBillDto 结构体
type FundBillDto struct {
	// 支付渠道代码
	FundChannel string `json:"fund_channel,omitempty" xml:"fund_channel,omitempty"`
	// 金额
	Amount string `json:"amount,omitempty" xml:"amount,omitempty"`
}
