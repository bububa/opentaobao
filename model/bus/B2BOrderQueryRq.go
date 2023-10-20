package bus

import (
	"sync"
)

// B2BOrderQueryRq 结构体
type B2BOrderQueryRq struct {
	// 阿里订单号
	AliTripOrderId string `json:"ali_trip_order_id,omitempty" xml:"ali_trip_order_id,omitempty"`
}

var poolB2BOrderQueryRq = sync.Pool{
	New: func() any {
		return new(B2BOrderQueryRq)
	},
}

// GetB2BOrderQueryRq() 从对象池中获取B2BOrderQueryRq
func GetB2BOrderQueryRq() *B2BOrderQueryRq {
	return poolB2BOrderQueryRq.Get().(*B2BOrderQueryRq)
}

// ReleaseB2BOrderQueryRq 释放B2BOrderQueryRq
func ReleaseB2BOrderQueryRq(v *B2BOrderQueryRq) {
	v.AliTripOrderId = ""
	poolB2BOrderQueryRq.Put(v)
}
