package btrip

import (
	"sync"
)

// BtripFlightCancelOrderRq 结构体
type BtripFlightCancelOrderRq struct {
	// 子渠道例如corpId
	SubChannel string `json:"sub_channel,omitempty" xml:"sub_channel,omitempty"`
	// 外部订单号
	DisOrderId string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
}

var poolBtripFlightCancelOrderRq = sync.Pool{
	New: func() any {
		return new(BtripFlightCancelOrderRq)
	},
}

// GetBtripFlightCancelOrderRq() 从对象池中获取BtripFlightCancelOrderRq
func GetBtripFlightCancelOrderRq() *BtripFlightCancelOrderRq {
	return poolBtripFlightCancelOrderRq.Get().(*BtripFlightCancelOrderRq)
}

// ReleaseBtripFlightCancelOrderRq 释放BtripFlightCancelOrderRq
func ReleaseBtripFlightCancelOrderRq(v *BtripFlightCancelOrderRq) {
	v.SubChannel = ""
	v.DisOrderId = ""
	poolBtripFlightCancelOrderRq.Put(v)
}
