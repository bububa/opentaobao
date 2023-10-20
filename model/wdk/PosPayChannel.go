package wdk

import (
	"sync"
)

// PosPayChannel 结构体
type PosPayChannel struct {
	// 支付方式编码，盒马给出了常见支付方式的编码
	PayType string `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	// 该支付方式对应的支付金额
	PayAmount int64 `json:"pay_amount,omitempty" xml:"pay_amount,omitempty"`
}

var poolPosPayChannel = sync.Pool{
	New: func() any {
		return new(PosPayChannel)
	},
}

// GetPosPayChannel() 从对象池中获取PosPayChannel
func GetPosPayChannel() *PosPayChannel {
	return poolPosPayChannel.Get().(*PosPayChannel)
}

// ReleasePosPayChannel 释放PosPayChannel
func ReleasePosPayChannel(v *PosPayChannel) {
	v.PayType = ""
	v.PayAmount = 0
	poolPosPayChannel.Put(v)
}
