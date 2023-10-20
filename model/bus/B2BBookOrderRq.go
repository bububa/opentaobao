package bus

import (
	"sync"
)

// B2BBookOrderRq 结构体
type B2BBookOrderRq struct {
	// 阿里订单号
	AliOrderId string `json:"ali_order_id,omitempty" xml:"ali_order_id,omitempty"`
}

var poolB2BBookOrderRq = sync.Pool{
	New: func() any {
		return new(B2BBookOrderRq)
	},
}

// GetB2BBookOrderRq() 从对象池中获取B2BBookOrderRq
func GetB2BBookOrderRq() *B2BBookOrderRq {
	return poolB2BBookOrderRq.Get().(*B2BBookOrderRq)
}

// ReleaseB2BBookOrderRq 释放B2BBookOrderRq
func ReleaseB2BBookOrderRq(v *B2BBookOrderRq) {
	v.AliOrderId = ""
	poolB2BBookOrderRq.Put(v)
}
