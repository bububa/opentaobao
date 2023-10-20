package btrip

import (
	"sync"
)

// BtripFlightModifyPayRq 结构体
type BtripFlightModifyPayRq struct {
	// 分销外部订单号
	DisOrderId string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
	// 分销子渠道，通常为corpId
	SubChannel string `json:"sub_channel,omitempty" xml:"sub_channel,omitempty"`
	// 分销外部改签单号
	DisSubOrderId string `json:"dis_sub_order_id,omitempty" xml:"dis_sub_order_id,omitempty"`
	// 额外信息
	Extra string `json:"extra,omitempty" xml:"extra,omitempty"`
	// 支付总金额
	TotalPayPrice int64 `json:"total_pay_price,omitempty" xml:"total_pay_price,omitempty"`
	// 企业支付金额
	CorpPayPrice int64 `json:"corp_pay_price,omitempty" xml:"corp_pay_price,omitempty"`
	// 个人支付金额
	PersonalPayPrice int64 `json:"personal_pay_price,omitempty" xml:"personal_pay_price,omitempty"`
}

var poolBtripFlightModifyPayRq = sync.Pool{
	New: func() any {
		return new(BtripFlightModifyPayRq)
	},
}

// GetBtripFlightModifyPayRq() 从对象池中获取BtripFlightModifyPayRq
func GetBtripFlightModifyPayRq() *BtripFlightModifyPayRq {
	return poolBtripFlightModifyPayRq.Get().(*BtripFlightModifyPayRq)
}

// ReleaseBtripFlightModifyPayRq 释放BtripFlightModifyPayRq
func ReleaseBtripFlightModifyPayRq(v *BtripFlightModifyPayRq) {
	v.DisOrderId = ""
	v.SubChannel = ""
	v.DisSubOrderId = ""
	v.Extra = ""
	v.TotalPayPrice = 0
	v.CorpPayPrice = 0
	v.PersonalPayPrice = 0
	poolBtripFlightModifyPayRq.Put(v)
}
