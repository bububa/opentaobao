package wdk

import (
	"sync"
)

// Outsuborders 结构体
type Outsuborders struct {
	// 渠道子订单ID
	OutSubOrderId string `json:"out_sub_order_id,omitempty" xml:"out_sub_order_id,omitempty"`
	// 渠道子订单最大可退金额
	MaxRefundFee int64 `json:"max_refund_fee,omitempty" xml:"max_refund_fee,omitempty"`
	// 当逆向已达到整单退时, 最后一个会有退运费金额
	PostFee int64 `json:"post_fee,omitempty" xml:"post_fee,omitempty"`
	// 是否可退
	CanReverse bool `json:"can_reverse,omitempty" xml:"can_reverse,omitempty"`
}

var poolOutsuborders = sync.Pool{
	New: func() any {
		return new(Outsuborders)
	},
}

// GetOutsuborders() 从对象池中获取Outsuborders
func GetOutsuborders() *Outsuborders {
	return poolOutsuborders.Get().(*Outsuborders)
}

// ReleaseOutsuborders 释放Outsuborders
func ReleaseOutsuborders(v *Outsuborders) {
	v.OutSubOrderId = ""
	v.MaxRefundFee = 0
	v.PostFee = 0
	v.CanReverse = false
	poolOutsuborders.Put(v)
}
