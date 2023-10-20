package tbrefund

import (
	"sync"
)

// MaxRefundFee 结构体
type MaxRefundFee struct {
	// 可以协商的最大退款金额
	MaxRefundFee int64 `json:"max_refund_fee,omitempty" xml:"max_refund_fee,omitempty"`
}

var poolMaxRefundFee = sync.Pool{
	New: func() any {
		return new(MaxRefundFee)
	},
}

// GetMaxRefundFee() 从对象池中获取MaxRefundFee
func GetMaxRefundFee() *MaxRefundFee {
	return poolMaxRefundFee.Get().(*MaxRefundFee)
}

// ReleaseMaxRefundFee 释放MaxRefundFee
func ReleaseMaxRefundFee(v *MaxRefundFee) {
	v.MaxRefundFee = 0
	poolMaxRefundFee.Put(v)
}
