package alicom

import (
	"sync"
)

// BankCreditCheckResponse 结构体
type BankCreditCheckResponse struct {
	// 能否申请
	CanApply bool `json:"can_apply,omitempty" xml:"can_apply,omitempty"`
}

var poolBankCreditCheckResponse = sync.Pool{
	New: func() any {
		return new(BankCreditCheckResponse)
	},
}

// GetBankCreditCheckResponse() 从对象池中获取BankCreditCheckResponse
func GetBankCreditCheckResponse() *BankCreditCheckResponse {
	return poolBankCreditCheckResponse.Get().(*BankCreditCheckResponse)
}

// ReleaseBankCreditCheckResponse 释放BankCreditCheckResponse
func ReleaseBankCreditCheckResponse(v *BankCreditCheckResponse) {
	v.CanApply = false
	poolBankCreditCheckResponse.Put(v)
}
