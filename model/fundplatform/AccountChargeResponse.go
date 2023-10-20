package fundplatform

import (
	"sync"
)

// AccountChargeResponse 结构体
type AccountChargeResponse struct {
	// 充值URL
	PayUrl string `json:"pay_url,omitempty" xml:"pay_url,omitempty"`
	// 充值的账户ID
	AccountId int64 `json:"account_id,omitempty" xml:"account_id,omitempty"`
}

var poolAccountChargeResponse = sync.Pool{
	New: func() any {
		return new(AccountChargeResponse)
	},
}

// GetAccountChargeResponse() 从对象池中获取AccountChargeResponse
func GetAccountChargeResponse() *AccountChargeResponse {
	return poolAccountChargeResponse.Get().(*AccountChargeResponse)
}

// ReleaseAccountChargeResponse 释放AccountChargeResponse
func ReleaseAccountChargeResponse(v *AccountChargeResponse) {
	v.PayUrl = ""
	v.AccountId = 0
	poolAccountChargeResponse.Put(v)
}
