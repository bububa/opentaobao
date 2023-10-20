package alicom

import (
	"sync"
)

// BankCreditCheckRequest 结构体
type BankCreditCheckRequest struct {
	// 银行的编码
	BankCode string `json:"bank_code,omitempty" xml:"bank_code,omitempty"`
	// 用户的openId
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
}

var poolBankCreditCheckRequest = sync.Pool{
	New: func() any {
		return new(BankCreditCheckRequest)
	},
}

// GetBankCreditCheckRequest() 从对象池中获取BankCreditCheckRequest
func GetBankCreditCheckRequest() *BankCreditCheckRequest {
	return poolBankCreditCheckRequest.Get().(*BankCreditCheckRequest)
}

// ReleaseBankCreditCheckRequest 释放BankCreditCheckRequest
func ReleaseBankCreditCheckRequest(v *BankCreditCheckRequest) {
	v.BankCode = ""
	v.OpenId = ""
	poolBankCreditCheckRequest.Put(v)
}
