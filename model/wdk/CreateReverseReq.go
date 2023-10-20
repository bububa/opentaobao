package wdk

import (
	"sync"
)

// CreateReverseReq 结构体
type CreateReverseReq struct {
	// wdk子单号
	BizOrderIds []int64 `json:"biz_order_ids,omitempty" xml:"biz_order_ids>int64,omitempty"`
	// 礼品卡号
	GiftCardNos []string `json:"gift_card_nos,omitempty" xml:"gift_card_nos>string,omitempty"`
	// 退款凭证
	Proofs []string `json:"proofs,omitempty" xml:"proofs>string,omitempty"`
	// 退款渠道
	RefundChannelList []RefundChannelVo `json:"refund_channel_list,omitempty" xml:"refund_channel_list>refund_channel_vo,omitempty"`
	// 退款原因描述
	ReasonText string `json:"reason_text,omitempty" xml:"reason_text,omitempty"`
	// 请求id（apply接口返回）
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 门店ID
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 操作人
	Operator *OperatorVo `json:"operator,omitempty" xml:"operator,omitempty"`
	// 退款原因id
	ReasonId int64 `json:"reason_id,omitempty" xml:"reason_id,omitempty"`
	// 退款金额
	RefundAmount int64 `json:"refund_amount,omitempty" xml:"refund_amount,omitempty"`
}

var poolCreateReverseReq = sync.Pool{
	New: func() any {
		return new(CreateReverseReq)
	},
}

// GetCreateReverseReq() 从对象池中获取CreateReverseReq
func GetCreateReverseReq() *CreateReverseReq {
	return poolCreateReverseReq.Get().(*CreateReverseReq)
}

// ReleaseCreateReverseReq 释放CreateReverseReq
func ReleaseCreateReverseReq(v *CreateReverseReq) {
	v.BizOrderIds = v.BizOrderIds[:0]
	v.GiftCardNos = v.GiftCardNos[:0]
	v.Proofs = v.Proofs[:0]
	v.RefundChannelList = v.RefundChannelList[:0]
	v.ReasonText = ""
	v.RequestId = ""
	v.StoreId = ""
	v.Operator = nil
	v.ReasonId = 0
	v.RefundAmount = 0
	poolCreateReverseReq.Put(v)
}
