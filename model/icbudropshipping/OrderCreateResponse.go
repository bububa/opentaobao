package icbudropshipping

import (
	"sync"
)

// OrderCreateResponse 结构体
type OrderCreateResponse struct {
	// pay url
	PayUrl string `json:"pay_url,omitempty" xml:"pay_url,omitempty"`
	// order number
	TradeId string `json:"trade_id,omitempty" xml:"trade_id,omitempty"`
}

var poolOrderCreateResponse = sync.Pool{
	New: func() any {
		return new(OrderCreateResponse)
	},
}

// GetOrderCreateResponse() 从对象池中获取OrderCreateResponse
func GetOrderCreateResponse() *OrderCreateResponse {
	return poolOrderCreateResponse.Get().(*OrderCreateResponse)
}

// ReleaseOrderCreateResponse 释放OrderCreateResponse
func ReleaseOrderCreateResponse(v *OrderCreateResponse) {
	v.PayUrl = ""
	v.TradeId = ""
	poolOrderCreateResponse.Put(v)
}
