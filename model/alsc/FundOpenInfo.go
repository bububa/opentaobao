package alsc

// FundOpenInfo 结构体
type FundOpenInfo struct {
	// 总金额
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 用户实付
	BuyerPayAmt int64 `json:"buyer_pay_amt,omitempty" xml:"buyer_pay_amt,omitempty"`
	// 商户补贴
	MerchantSubsidy int64 `json:"merchant_subsidy,omitempty" xml:"merchant_subsidy,omitempty"`
	// 平台补贴
	PlatformSubsidy int64 `json:"platform_subsidy,omitempty" xml:"platform_subsidy,omitempty"`
}
