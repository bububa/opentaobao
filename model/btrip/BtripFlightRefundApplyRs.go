package btrip

import (
	"sync"
)

// BtripFlightRefundApplyRs 结构体
type BtripFlightRefundApplyRs struct {
	// 分销外部订单号
	DisOrderId string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
	// 分销外部退票单号
	DisSubOrderId string `json:"dis_sub_order_id,omitempty" xml:"dis_sub_order_id,omitempty"`
	// 退票手续费
	RefundFee int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 退票应退
	RefundMoney int64 `json:"refund_money,omitempty" xml:"refund_money,omitempty"`
	// 商旅退票申请单号
	RefundApplyId int64 `json:"refund_apply_id,omitempty" xml:"refund_apply_id,omitempty"`
}

var poolBtripFlightRefundApplyRs = sync.Pool{
	New: func() any {
		return new(BtripFlightRefundApplyRs)
	},
}

// GetBtripFlightRefundApplyRs() 从对象池中获取BtripFlightRefundApplyRs
func GetBtripFlightRefundApplyRs() *BtripFlightRefundApplyRs {
	return poolBtripFlightRefundApplyRs.Get().(*BtripFlightRefundApplyRs)
}

// ReleaseBtripFlightRefundApplyRs 释放BtripFlightRefundApplyRs
func ReleaseBtripFlightRefundApplyRs(v *BtripFlightRefundApplyRs) {
	v.DisOrderId = ""
	v.DisSubOrderId = ""
	v.RefundFee = 0
	v.RefundMoney = 0
	v.RefundApplyId = 0
	poolBtripFlightRefundApplyRs.Put(v)
}
