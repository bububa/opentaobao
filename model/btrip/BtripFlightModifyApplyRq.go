package btrip

import (
	"sync"
)

// BtripFlightModifyApplyRq 结构体
type BtripFlightModifyApplyRq struct {
	// 改签航班信息
	ModifyFlightInfoList []ModifyFlightInfo `json:"modify_flight_info_list,omitempty" xml:"modify_flight_info_list>modify_flight_info,omitempty"`
	// 分销外部订单号
	DisOrderId string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
	// 改签原因
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 分销子渠道，通常为corpId
	SubChannel string `json:"sub_channel,omitempty" xml:"sub_channel,omitempty"`
	// 分销外部改签单号
	DisSubOrderId string `json:"dis_sub_order_id,omitempty" xml:"dis_sub_order_id,omitempty"`
	// 联系人电话
	ContactPhone string `json:"contact_phone,omitempty" xml:"contact_phone,omitempty"`
	// ota商品itemId
	OtaItemId string `json:"ota_item_id,omitempty" xml:"ota_item_id,omitempty"`
	// 会话ID
	SessionId string `json:"session_id,omitempty" xml:"session_id,omitempty"`
	// 是否自愿
	IsVoluntary int64 `json:"is_voluntary,omitempty" xml:"is_voluntary,omitempty"`
	// 默认为false
	WhetherRetry bool `json:"whether_retry,omitempty" xml:"whether_retry,omitempty"`
}

var poolBtripFlightModifyApplyRq = sync.Pool{
	New: func() any {
		return new(BtripFlightModifyApplyRq)
	},
}

// GetBtripFlightModifyApplyRq() 从对象池中获取BtripFlightModifyApplyRq
func GetBtripFlightModifyApplyRq() *BtripFlightModifyApplyRq {
	return poolBtripFlightModifyApplyRq.Get().(*BtripFlightModifyApplyRq)
}

// ReleaseBtripFlightModifyApplyRq 释放BtripFlightModifyApplyRq
func ReleaseBtripFlightModifyApplyRq(v *BtripFlightModifyApplyRq) {
	v.ModifyFlightInfoList = v.ModifyFlightInfoList[:0]
	v.DisOrderId = ""
	v.Reason = ""
	v.SubChannel = ""
	v.DisSubOrderId = ""
	v.ContactPhone = ""
	v.OtaItemId = ""
	v.SessionId = ""
	v.IsVoluntary = 0
	v.WhetherRetry = false
	poolBtripFlightModifyApplyRq.Put(v)
}
