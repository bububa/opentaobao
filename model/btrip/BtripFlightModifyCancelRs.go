package btrip

import (
	"sync"
)

// BtripFlightModifyCancelRs 结构体
type BtripFlightModifyCancelRs struct {
	// 外部分销改签订单号
	DisSubOrderId string `json:"dis_sub_order_id,omitempty" xml:"dis_sub_order_id,omitempty"`
	// 改签取消时间
	CancelTime string `json:"cancel_time,omitempty" xml:"cancel_time,omitempty"`
	// 改签单的状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
}

var poolBtripFlightModifyCancelRs = sync.Pool{
	New: func() any {
		return new(BtripFlightModifyCancelRs)
	},
}

// GetBtripFlightModifyCancelRs() 从对象池中获取BtripFlightModifyCancelRs
func GetBtripFlightModifyCancelRs() *BtripFlightModifyCancelRs {
	return poolBtripFlightModifyCancelRs.Get().(*BtripFlightModifyCancelRs)
}

// ReleaseBtripFlightModifyCancelRs 释放BtripFlightModifyCancelRs
func ReleaseBtripFlightModifyCancelRs(v *BtripFlightModifyCancelRs) {
	v.DisSubOrderId = ""
	v.CancelTime = ""
	v.Status = ""
	poolBtripFlightModifyCancelRs.Put(v)
}
