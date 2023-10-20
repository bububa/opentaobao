package btrip

import (
	"sync"
)

// BtripFlightRefundPreCalRs 结构体
type BtripFlightRefundPreCalRs struct {
	// 多行程退票预计算list
	MultiRefundPreCalDetailList []MultiRefundPreCalDetail `json:"multi_refund_pre_cal_detail_list,omitempty" xml:"multi_refund_pre_cal_detail_list>multi_refund_pre_cal_detail,omitempty"`
	// 退票原因
	ReturnReason []ReturnReasonDetail `json:"return_reason,omitempty" xml:"return_reason>return_reason_detail,omitempty"`
	// 会话id
	SessionId string `json:"session_id,omitempty" xml:"session_id,omitempty"`
	// 商品id
	ItemUnitId string `json:"item_unit_id,omitempty" xml:"item_unit_id,omitempty"`
	// 提示
	Tips string `json:"tips,omitempty" xml:"tips,omitempty"`
	// 预计退还金额
	PreRefundMoney int64 `json:"pre_refund_money,omitempty" xml:"pre_refund_money,omitempty"`
	// 退票手续费
	RefundFee int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 是否发生航变
	FlightChange bool `json:"flight_change,omitempty" xml:"flight_change,omitempty"`
}

var poolBtripFlightRefundPreCalRs = sync.Pool{
	New: func() any {
		return new(BtripFlightRefundPreCalRs)
	},
}

// GetBtripFlightRefundPreCalRs() 从对象池中获取BtripFlightRefundPreCalRs
func GetBtripFlightRefundPreCalRs() *BtripFlightRefundPreCalRs {
	return poolBtripFlightRefundPreCalRs.Get().(*BtripFlightRefundPreCalRs)
}

// ReleaseBtripFlightRefundPreCalRs 释放BtripFlightRefundPreCalRs
func ReleaseBtripFlightRefundPreCalRs(v *BtripFlightRefundPreCalRs) {
	v.MultiRefundPreCalDetailList = v.MultiRefundPreCalDetailList[:0]
	v.ReturnReason = v.ReturnReason[:0]
	v.SessionId = ""
	v.ItemUnitId = ""
	v.Tips = ""
	v.PreRefundMoney = 0
	v.RefundFee = 0
	v.FlightChange = false
	poolBtripFlightRefundPreCalRs.Put(v)
}
