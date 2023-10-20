package wdk

import (
	"sync"
)

// OrderPayChannel 结构体
type OrderPayChannel struct {
	// 支付方式
	PayType string `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	// 支付渠道名称
	PayChannel string `json:"pay_channel,omitempty" xml:"pay_channel,omitempty"`
	// 支付金额
	PayFee int64 `json:"pay_fee,omitempty" xml:"pay_fee,omitempty"`
}

var poolOrderPayChannel = sync.Pool{
	New: func() any {
		return new(OrderPayChannel)
	},
}

// GetOrderPayChannel() 从对象池中获取OrderPayChannel
func GetOrderPayChannel() *OrderPayChannel {
	return poolOrderPayChannel.Get().(*OrderPayChannel)
}

// ReleaseOrderPayChannel 释放OrderPayChannel
func ReleaseOrderPayChannel(v *OrderPayChannel) {
	v.PayType = ""
	v.PayChannel = ""
	v.PayFee = 0
	poolOrderPayChannel.Put(v)
}
