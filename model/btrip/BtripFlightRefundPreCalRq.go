package btrip

import (
	"sync"
)

// BtripFlightRefundPreCalRq 结构体
type BtripFlightRefundPreCalRq struct {
	// 乘客航段信息
	PassengerSegmentInfoList []PassengerSegmentInfo `json:"passenger_segment_info_list,omitempty" xml:"passenger_segment_info_list>passenger_segment_info,omitempty"`
	// 退票票号列表
	TicketNos []string `json:"ticket_nos,omitempty" xml:"ticket_nos>string,omitempty"`
	// 分销外部订单号
	DisOrderId string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
	// 子渠道，通常为corpId
	SubChannel string `json:"sub_channel,omitempty" xml:"sub_channel,omitempty"`
	// 是否自愿
	IsVoluntary int64 `json:"is_voluntary,omitempty" xml:"is_voluntary,omitempty"`
}

var poolBtripFlightRefundPreCalRq = sync.Pool{
	New: func() any {
		return new(BtripFlightRefundPreCalRq)
	},
}

// GetBtripFlightRefundPreCalRq() 从对象池中获取BtripFlightRefundPreCalRq
func GetBtripFlightRefundPreCalRq() *BtripFlightRefundPreCalRq {
	return poolBtripFlightRefundPreCalRq.Get().(*BtripFlightRefundPreCalRq)
}

// ReleaseBtripFlightRefundPreCalRq 释放BtripFlightRefundPreCalRq
func ReleaseBtripFlightRefundPreCalRq(v *BtripFlightRefundPreCalRq) {
	v.PassengerSegmentInfoList = v.PassengerSegmentInfoList[:0]
	v.TicketNos = v.TicketNos[:0]
	v.DisOrderId = ""
	v.SubChannel = ""
	v.IsVoluntary = 0
	poolBtripFlightRefundPreCalRq.Put(v)
}
