package alicom

import (
	"sync"
)

// QueryOrderReq 结构体
type QueryOrderReq struct {
	// 外部订单id
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
}

var poolQueryOrderReq = sync.Pool{
	New: func() any {
		return new(QueryOrderReq)
	},
}

// GetQueryOrderReq() 从对象池中获取QueryOrderReq
func GetQueryOrderReq() *QueryOrderReq {
	return poolQueryOrderReq.Get().(*QueryOrderReq)
}

// ReleaseQueryOrderReq 释放QueryOrderReq
func ReleaseQueryOrderReq(v *QueryOrderReq) {
	v.OutOrderId = ""
	poolQueryOrderReq.Put(v)
}
