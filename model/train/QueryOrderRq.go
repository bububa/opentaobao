package train

import (
	"sync"
)

// QueryOrderRq 结构体
type QueryOrderRq struct {
	// 主单号
	TtpOrderId int64 `json:"ttp_order_id,omitempty" xml:"ttp_order_id,omitempty"`
}

var poolQueryOrderRq = sync.Pool{
	New: func() any {
		return new(QueryOrderRq)
	},
}

// GetQueryOrderRq() 从对象池中获取QueryOrderRq
func GetQueryOrderRq() *QueryOrderRq {
	return poolQueryOrderRq.Get().(*QueryOrderRq)
}

// ReleaseQueryOrderRq 释放QueryOrderRq
func ReleaseQueryOrderRq(v *QueryOrderRq) {
	v.TtpOrderId = 0
	poolQueryOrderRq.Put(v)
}
