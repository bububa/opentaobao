package train

import (
	"sync"
)

// IgnoreOrderRq 结构体
type IgnoreOrderRq struct {
	// 子单号 （不填默认全部忽略）
	SubOrderId []int64 `json:"sub_order_id,omitempty" xml:"sub_order_id>int64,omitempty"`
	// 主单号
	TtpOrderId int64 `json:"ttp_order_id,omitempty" xml:"ttp_order_id,omitempty"`
}

var poolIgnoreOrderRq = sync.Pool{
	New: func() any {
		return new(IgnoreOrderRq)
	},
}

// GetIgnoreOrderRq() 从对象池中获取IgnoreOrderRq
func GetIgnoreOrderRq() *IgnoreOrderRq {
	return poolIgnoreOrderRq.Get().(*IgnoreOrderRq)
}

// ReleaseIgnoreOrderRq 释放IgnoreOrderRq
func ReleaseIgnoreOrderRq(v *IgnoreOrderRq) {
	v.SubOrderId = v.SubOrderId[:0]
	v.TtpOrderId = 0
	poolIgnoreOrderRq.Put(v)
}
