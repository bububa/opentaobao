package trade

import (
	"sync"
)

// CheckoutExchangeRequest 结构体
type CheckoutExchangeRequest struct {
	// 基准币种(卖家设置的）
	BaseCurrency string `json:"base_currency,omitempty" xml:"base_currency,omitempty"`
	// 报价币种(买家看到的）
	QuoteCurrency string `json:"quote_currency,omitempty" xml:"quote_currency,omitempty"`
}

var poolCheckoutExchangeRequest = sync.Pool{
	New: func() any {
		return new(CheckoutExchangeRequest)
	},
}

// GetCheckoutExchangeRequest() 从对象池中获取CheckoutExchangeRequest
func GetCheckoutExchangeRequest() *CheckoutExchangeRequest {
	return poolCheckoutExchangeRequest.Get().(*CheckoutExchangeRequest)
}

// ReleaseCheckoutExchangeRequest 释放CheckoutExchangeRequest
func ReleaseCheckoutExchangeRequest(v *CheckoutExchangeRequest) {
	v.BaseCurrency = ""
	v.QuoteCurrency = ""
	poolCheckoutExchangeRequest.Put(v)
}
