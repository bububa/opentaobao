package drug

import (
	"sync"
)

// FundBillDto 结构体
type FundBillDto struct {
	// 支付渠道代码
	FundChannel string `json:"fund_channel,omitempty" xml:"fund_channel,omitempty"`
	// 金额
	Amount string `json:"amount,omitempty" xml:"amount,omitempty"`
}

var poolFundBillDto = sync.Pool{
	New: func() any {
		return new(FundBillDto)
	},
}

// GetFundBillDto() 从对象池中获取FundBillDto
func GetFundBillDto() *FundBillDto {
	return poolFundBillDto.Get().(*FundBillDto)
}

// ReleaseFundBillDto 释放FundBillDto
func ReleaseFundBillDto(v *FundBillDto) {
	v.FundChannel = ""
	v.Amount = ""
	poolFundBillDto.Put(v)
}
