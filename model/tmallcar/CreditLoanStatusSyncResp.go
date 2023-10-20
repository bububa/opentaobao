package tmallcar

import (
	"sync"
)

// CreditLoanStatusSyncResp 结构体
type CreditLoanStatusSyncResp struct {
	// 是否成功
	Succeed bool `json:"succeed,omitempty" xml:"succeed,omitempty"`
	// 是否重试
	Retry bool `json:"retry,omitempty" xml:"retry,omitempty"`
}

var poolCreditLoanStatusSyncResp = sync.Pool{
	New: func() any {
		return new(CreditLoanStatusSyncResp)
	},
}

// GetCreditLoanStatusSyncResp() 从对象池中获取CreditLoanStatusSyncResp
func GetCreditLoanStatusSyncResp() *CreditLoanStatusSyncResp {
	return poolCreditLoanStatusSyncResp.Get().(*CreditLoanStatusSyncResp)
}

// ReleaseCreditLoanStatusSyncResp 释放CreditLoanStatusSyncResp
func ReleaseCreditLoanStatusSyncResp(v *CreditLoanStatusSyncResp) {
	v.Succeed = false
	v.Retry = false
	poolCreditLoanStatusSyncResp.Put(v)
}
