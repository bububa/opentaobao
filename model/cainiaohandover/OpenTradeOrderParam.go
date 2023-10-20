package cainiaohandover

import (
	"sync"
)

// OpenTradeOrderParam 结构体
type OpenTradeOrderParam struct {
	// 主交易单ID
	TradeOrderId int64 `json:"trade_order_id,omitempty" xml:"trade_order_id,omitempty"`
}

var poolOpenTradeOrderParam = sync.Pool{
	New: func() any {
		return new(OpenTradeOrderParam)
	},
}

// GetOpenTradeOrderParam() 从对象池中获取OpenTradeOrderParam
func GetOpenTradeOrderParam() *OpenTradeOrderParam {
	return poolOpenTradeOrderParam.Get().(*OpenTradeOrderParam)
}

// ReleaseOpenTradeOrderParam 释放OpenTradeOrderParam
func ReleaseOpenTradeOrderParam(v *OpenTradeOrderParam) {
	v.TradeOrderId = 0
	poolOpenTradeOrderParam.Put(v)
}
