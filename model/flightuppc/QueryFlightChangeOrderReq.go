package flightuppc

import (
	"sync"
)

// QueryFlightChangeOrderReq 结构体
type QueryFlightChangeOrderReq struct {
	// 订单号
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolQueryFlightChangeOrderReq = sync.Pool{
	New: func() any {
		return new(QueryFlightChangeOrderReq)
	},
}

// GetQueryFlightChangeOrderReq() 从对象池中获取QueryFlightChangeOrderReq
func GetQueryFlightChangeOrderReq() *QueryFlightChangeOrderReq {
	return poolQueryFlightChangeOrderReq.Get().(*QueryFlightChangeOrderReq)
}

// ReleaseQueryFlightChangeOrderReq 释放QueryFlightChangeOrderReq
func ReleaseQueryFlightChangeOrderReq(v *QueryFlightChangeOrderReq) {
	v.OrderId = 0
	poolQueryFlightChangeOrderReq.Put(v)
}
