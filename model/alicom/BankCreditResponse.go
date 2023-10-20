package alicom

import (
	"sync"
)

// BankCreditResponse 结构体
type BankCreditResponse struct {
	// 银行编码
	BankCode string `json:"bank_code,omitempty" xml:"bank_code,omitempty"`
	// 用户在该行的申请状态
	State string `json:"state,omitempty" xml:"state,omitempty"`
	// 事件发生的时间戳
	EventTime int64 `json:"event_time,omitempty" xml:"event_time,omitempty"`
}

var poolBankCreditResponse = sync.Pool{
	New: func() any {
		return new(BankCreditResponse)
	},
}

// GetBankCreditResponse() 从对象池中获取BankCreditResponse
func GetBankCreditResponse() *BankCreditResponse {
	return poolBankCreditResponse.Get().(*BankCreditResponse)
}

// ReleaseBankCreditResponse 释放BankCreditResponse
func ReleaseBankCreditResponse(v *BankCreditResponse) {
	v.BankCode = ""
	v.State = ""
	v.EventTime = 0
	poolBankCreditResponse.Put(v)
}
