package btrip

import (
	"sync"
)

// BtripFlightModifyCancelRq 结构体
type BtripFlightModifyCancelRq struct {
	// 分销外部订单号
	DisSubOrderId string `json:"dis_sub_order_id,omitempty" xml:"dis_sub_order_id,omitempty"`
	// 分销外部改签单号
	DisOrderId string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
	// 分销子渠道，通常为商旅corpId
	SubChannel string `json:"sub_channel,omitempty" xml:"sub_channel,omitempty"`
}

var poolBtripFlightModifyCancelRq = sync.Pool{
	New: func() any {
		return new(BtripFlightModifyCancelRq)
	},
}

// GetBtripFlightModifyCancelRq() 从对象池中获取BtripFlightModifyCancelRq
func GetBtripFlightModifyCancelRq() *BtripFlightModifyCancelRq {
	return poolBtripFlightModifyCancelRq.Get().(*BtripFlightModifyCancelRq)
}

// ReleaseBtripFlightModifyCancelRq 释放BtripFlightModifyCancelRq
func ReleaseBtripFlightModifyCancelRq(v *BtripFlightModifyCancelRq) {
	v.DisSubOrderId = ""
	v.DisOrderId = ""
	v.SubChannel = ""
	poolBtripFlightModifyCancelRq.Put(v)
}
