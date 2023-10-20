package wdk

import (
	"sync"
)

// ApplyReverseResponse 结构体
type ApplyReverseResponse struct {
	// wdk交易单号
	BizOrderIds []int64 `json:"biz_order_ids,omitempty" xml:"biz_order_ids>int64,omitempty"`
	// 礼品卡号
	GiftCardNos []string `json:"gift_card_nos,omitempty" xml:"gift_card_nos>string,omitempty"`
	// 原因
	ReasonList []ReasonVo `json:"reason_list,omitempty" xml:"reason_list>reason_vo,omitempty"`
	// 退款渠道
	RefundChannelList []RefundChannelVo `json:"refund_channel_list,omitempty" xml:"refund_channel_list>refund_channel_vo,omitempty"`
	// 逆向单id
	ReverseIds []int64 `json:"reverse_ids,omitempty" xml:"reverse_ids>int64,omitempty"`
	// 请求id（提交申请入参）
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 门店id
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 1售中 2售后
	InSaleRefund int64 `json:"in_sale_refund,omitempty" xml:"in_sale_refund,omitempty"`
	// 最大可退金额
	MaxRefundFee int64 `json:"max_refund_fee,omitempty" xml:"max_refund_fee,omitempty"`
	// 运费
	PostFee int64 `json:"post_fee,omitempty" xml:"post_fee,omitempty"`
	// 是否支持修改金额
	SupportModifyAmount bool `json:"support_modify_amount,omitempty" xml:"support_modify_amount,omitempty"`
}

var poolApplyReverseResponse = sync.Pool{
	New: func() any {
		return new(ApplyReverseResponse)
	},
}

// GetApplyReverseResponse() 从对象池中获取ApplyReverseResponse
func GetApplyReverseResponse() *ApplyReverseResponse {
	return poolApplyReverseResponse.Get().(*ApplyReverseResponse)
}

// ReleaseApplyReverseResponse 释放ApplyReverseResponse
func ReleaseApplyReverseResponse(v *ApplyReverseResponse) {
	v.BizOrderIds = v.BizOrderIds[:0]
	v.GiftCardNos = v.GiftCardNos[:0]
	v.ReasonList = v.ReasonList[:0]
	v.RefundChannelList = v.RefundChannelList[:0]
	v.ReverseIds = v.ReverseIds[:0]
	v.RequestId = ""
	v.StoreId = ""
	v.InSaleRefund = 0
	v.MaxRefundFee = 0
	v.PostFee = 0
	v.SupportModifyAmount = false
	poolApplyReverseResponse.Put(v)
}
