package icbudropshipping

import (
	"sync"
)

// CashierPayResponse 结构体
type CashierPayResponse struct {
	// Payment url
	PayUrl string `json:"pay_url,omitempty" xml:"pay_url,omitempty"`
	// Payment failed reason code
	ReasonCode string `json:"reason_code,omitempty" xml:"reason_code,omitempty"`
	// Payment failed reason message
	ReasonMessage string `json:"reason_message,omitempty" xml:"reason_message,omitempty"`
	// UNPAY Unpaid order &lt;br /&gt; PAYING The order is being paid and it needs to wait for about 1 minute,&lt;br /&gt; PAY_SUCCESS Order payment is successful,&lt;br /&gt; PAY_FAILED Order payment failed
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// order id
	TradeId int64 `json:"trade_id,omitempty" xml:"trade_id,omitempty"`
}

var poolCashierPayResponse = sync.Pool{
	New: func() any {
		return new(CashierPayResponse)
	},
}

// GetCashierPayResponse() 从对象池中获取CashierPayResponse
func GetCashierPayResponse() *CashierPayResponse {
	return poolCashierPayResponse.Get().(*CashierPayResponse)
}

// ReleaseCashierPayResponse 释放CashierPayResponse
func ReleaseCashierPayResponse(v *CashierPayResponse) {
	v.PayUrl = ""
	v.ReasonCode = ""
	v.ReasonMessage = ""
	v.Status = ""
	v.TradeId = 0
	poolCashierPayResponse.Put(v)
}
