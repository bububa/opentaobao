package alsc

import (
	"sync"
)

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

var poolFundOpenInfo = sync.Pool{
	New: func() any {
		return new(FundOpenInfo)
	},
}

// GetFundOpenInfo() 从对象池中获取FundOpenInfo
func GetFundOpenInfo() *FundOpenInfo {
	return poolFundOpenInfo.Get().(*FundOpenInfo)
}

// ReleaseFundOpenInfo 释放FundOpenInfo
func ReleaseFundOpenInfo(v *FundOpenInfo) {
	v.Amount = 0
	v.BuyerPayAmt = 0
	v.MerchantSubsidy = 0
	v.PlatformSubsidy = 0
	poolFundOpenInfo.Put(v)
}
