package wdk

import (
	"sync"
)

// SavePurchasePriceResponse 结构体
type SavePurchasePriceResponse struct {
	// 新创建的变价单id
	TicketId string `json:"ticket_id,omitempty" xml:"ticket_id,omitempty"`
}

var poolSavePurchasePriceResponse = sync.Pool{
	New: func() any {
		return new(SavePurchasePriceResponse)
	},
}

// GetSavePurchasePriceResponse() 从对象池中获取SavePurchasePriceResponse
func GetSavePurchasePriceResponse() *SavePurchasePriceResponse {
	return poolSavePurchasePriceResponse.Get().(*SavePurchasePriceResponse)
}

// ReleaseSavePurchasePriceResponse 释放SavePurchasePriceResponse
func ReleaseSavePurchasePriceResponse(v *SavePurchasePriceResponse) {
	v.TicketId = ""
	poolSavePurchasePriceResponse.Put(v)
}
