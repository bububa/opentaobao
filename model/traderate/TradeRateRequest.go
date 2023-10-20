package traderate

import (
	"sync"
)

// TradeRateRequest 结构体
type TradeRateRequest struct {
	// 评价创建时间,格式:yyyy-MM-dd HH:mm:ss
	Created string `json:"created,omitempty" xml:"created,omitempty"`
	// 交易ID
	Tid int64 `json:"tid,omitempty" xml:"tid,omitempty"`
	// 子订单ID
	Oid int64 `json:"oid,omitempty" xml:"oid,omitempty"`
}

var poolTradeRateRequest = sync.Pool{
	New: func() any {
		return new(TradeRateRequest)
	},
}

// GetTradeRateRequest() 从对象池中获取TradeRateRequest
func GetTradeRateRequest() *TradeRateRequest {
	return poolTradeRateRequest.Get().(*TradeRateRequest)
}

// ReleaseTradeRateRequest 释放TradeRateRequest
func ReleaseTradeRateRequest(v *TradeRateRequest) {
	v.Created = ""
	v.Tid = 0
	v.Oid = 0
	poolTradeRateRequest.Put(v)
}
