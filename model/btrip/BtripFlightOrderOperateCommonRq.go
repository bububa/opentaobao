package btrip

import (
	"sync"
)

// BtripFlightOrderOperateCommonRq 结构体
type BtripFlightOrderOperateCommonRq struct {
	// 分销外部订单号
	DisOrderId string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
	// 分销子渠道，通常为corpId
	SubChannel string `json:"sub_channel,omitempty" xml:"sub_channel,omitempty"`
	// 分销外部改签单号
	DisSubOrderId string `json:"dis_sub_order_id,omitempty" xml:"dis_sub_order_id,omitempty"`
}

var poolBtripFlightOrderOperateCommonRq = sync.Pool{
	New: func() any {
		return new(BtripFlightOrderOperateCommonRq)
	},
}

// GetBtripFlightOrderOperateCommonRq() 从对象池中获取BtripFlightOrderOperateCommonRq
func GetBtripFlightOrderOperateCommonRq() *BtripFlightOrderOperateCommonRq {
	return poolBtripFlightOrderOperateCommonRq.Get().(*BtripFlightOrderOperateCommonRq)
}

// ReleaseBtripFlightOrderOperateCommonRq 释放BtripFlightOrderOperateCommonRq
func ReleaseBtripFlightOrderOperateCommonRq(v *BtripFlightOrderOperateCommonRq) {
	v.DisOrderId = ""
	v.SubChannel = ""
	v.DisSubOrderId = ""
	poolBtripFlightOrderOperateCommonRq.Put(v)
}
