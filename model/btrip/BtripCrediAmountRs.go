package btrip

import (
	"sync"
)

// BtripCrediAmountRs 结构体
type BtripCrediAmountRs struct {
	// 已使用额度，分
	UsedCreditLimit string `json:"used_credit_limit,omitempty" xml:"used_credit_limit,omitempty"`
	// 激活/冻结
	StatusDesc string `json:"status_desc,omitempty" xml:"status_desc,omitempty"`
	// 月结额度
	CreditLimit int64 `json:"credit_limit,omitempty" xml:"credit_limit,omitempty"`
	// 冻结金额    分
	FreezeAmount int64 `json:"freeze_amount,omitempty" xml:"freeze_amount,omitempty"`
	// 已使用额度  分
	CreditBalance int64 `json:"credit_balance,omitempty" xml:"credit_balance,omitempty"`
	// 剩余可用余额  分
	AvailableAmount int64 `json:"available_amount,omitempty" xml:"available_amount,omitempty"`
	// 账户状态1:激活 0：冻结
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

var poolBtripCrediAmountRs = sync.Pool{
	New: func() any {
		return new(BtripCrediAmountRs)
	},
}

// GetBtripCrediAmountRs() 从对象池中获取BtripCrediAmountRs
func GetBtripCrediAmountRs() *BtripCrediAmountRs {
	return poolBtripCrediAmountRs.Get().(*BtripCrediAmountRs)
}

// ReleaseBtripCrediAmountRs 释放BtripCrediAmountRs
func ReleaseBtripCrediAmountRs(v *BtripCrediAmountRs) {
	v.UsedCreditLimit = ""
	v.StatusDesc = ""
	v.CreditLimit = 0
	v.FreezeAmount = 0
	v.CreditBalance = 0
	v.AvailableAmount = 0
	v.Status = 0
	poolBtripCrediAmountRs.Put(v)
}
