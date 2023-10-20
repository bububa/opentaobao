package wdk

import (
	"sync"
)

// BuyGiftActivityQueryRequest 结构体
type BuyGiftActivityQueryRequest struct {
	// erp外部活动id
	OutActId string `json:"out_act_id,omitempty" xml:"out_act_id,omitempty"`
	// 活动ID
	ActId int64 `json:"act_id,omitempty" xml:"act_id,omitempty"`
}

var poolBuyGiftActivityQueryRequest = sync.Pool{
	New: func() any {
		return new(BuyGiftActivityQueryRequest)
	},
}

// GetBuyGiftActivityQueryRequest() 从对象池中获取BuyGiftActivityQueryRequest
func GetBuyGiftActivityQueryRequest() *BuyGiftActivityQueryRequest {
	return poolBuyGiftActivityQueryRequest.Get().(*BuyGiftActivityQueryRequest)
}

// ReleaseBuyGiftActivityQueryRequest 释放BuyGiftActivityQueryRequest
func ReleaseBuyGiftActivityQueryRequest(v *BuyGiftActivityQueryRequest) {
	v.OutActId = ""
	v.ActId = 0
	poolBuyGiftActivityQueryRequest.Put(v)
}
