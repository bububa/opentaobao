package wdk

import (
	"sync"
)

// LoadReverseResponse 结构体
type LoadReverseResponse struct {
	// wdk单号list
	BizOrderIds []int64 `json:"biz_order_ids,omitempty" xml:"biz_order_ids>int64,omitempty"`
	// 外部单号list
	OutBizOrderIds []string `json:"out_biz_order_ids,omitempty" xml:"out_biz_order_ids>string,omitempty"`
	// 凭证图片
	Proofs []string `json:"proofs,omitempty" xml:"proofs>string,omitempty"`
	// 原因标签
	ReasonTags []string `json:"reason_tags,omitempty" xml:"reason_tags>string,omitempty"`
	// RefundChannelVo
	RefundChannelList []RefundChannelVo `json:"refund_channel_list,omitempty" xml:"refund_channel_list>refund_channel_vo,omitempty"`
	// 退款单ids
	ReverseIds []int64 `json:"reverse_ids,omitempty" xml:"reverse_ids>int64,omitempty"`
	// 发起时间
	CreateDate string `json:"create_date,omitempty" xml:"create_date,omitempty"`
	// 发起备注
	CreateMemo string `json:"create_memo,omitempty" xml:"create_memo,omitempty"`
	// 结束时间
	EndDate string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// 最终原因
	EndReason string `json:"end_reason,omitempty" xml:"end_reason,omitempty"`
	// 终结人
	Ender string `json:"ender,omitempty" xml:"ender,omitempty"`
	// 外部主单号
	MainOutOrderId string `json:"main_out_order_id,omitempty" xml:"main_out_order_id,omitempty"`
	// 修正备注
	ModifiedMemo string `json:"modified_memo,omitempty" xml:"modified_memo,omitempty"`
	// 修正原因描述
	ModifiedReasonText string `json:"modified_reason_text,omitempty" xml:"modified_reason_text,omitempty"`
	// 外部单号
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 发起原因描述
	ReasonText string `json:"reason_text,omitempty" xml:"reason_text,omitempty"`
	// 请求id
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 退款单状态
	ReverseStatusStr string `json:"reverse_status_str,omitempty" xml:"reverse_status_str,omitempty"`
	// 退款类型
	ReverseTypeStr string `json:"reverse_type_str,omitempty" xml:"reverse_type_str,omitempty"`
	// 门店id
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// wdk单号
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 退款渠道
	CreateChannel int64 `json:"create_channel,omitempty" xml:"create_channel,omitempty"`
	// 退货取货单ID
	FetchOrderId int64 `json:"fetch_order_id,omitempty" xml:"fetch_order_id,omitempty"`
	// wdk主单号
	MainBizOrderId int64 `json:"main_biz_order_id,omitempty" xml:"main_biz_order_id,omitempty"`
	// 修正原因ID
	ModifiedReasonId int64 `json:"modified_reason_id,omitempty" xml:"modified_reason_id,omitempty"`
	// 发起原因ID
	ReasonId int64 `json:"reason_id,omitempty" xml:"reason_id,omitempty"`
	// 退款金额
	RefundAmount int64 `json:"refund_amount,omitempty" xml:"refund_amount,omitempty"`
	// 退款单id
	ReverseId int64 `json:"reverse_id,omitempty" xml:"reverse_id,omitempty"`
	// 退款单状态
	ReverseStatus int64 `json:"reverse_status,omitempty" xml:"reverse_status,omitempty"`
	// 退款类型
	ReverseType int64 `json:"reverse_type,omitempty" xml:"reverse_type,omitempty"`
	// 卖家id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 是否称重品
	Weight bool `json:"weight,omitempty" xml:"weight,omitempty"`
}

var poolLoadReverseResponse = sync.Pool{
	New: func() any {
		return new(LoadReverseResponse)
	},
}

// GetLoadReverseResponse() 从对象池中获取LoadReverseResponse
func GetLoadReverseResponse() *LoadReverseResponse {
	return poolLoadReverseResponse.Get().(*LoadReverseResponse)
}

// ReleaseLoadReverseResponse 释放LoadReverseResponse
func ReleaseLoadReverseResponse(v *LoadReverseResponse) {
	v.BizOrderIds = v.BizOrderIds[:0]
	v.OutBizOrderIds = v.OutBizOrderIds[:0]
	v.Proofs = v.Proofs[:0]
	v.ReasonTags = v.ReasonTags[:0]
	v.RefundChannelList = v.RefundChannelList[:0]
	v.ReverseIds = v.ReverseIds[:0]
	v.CreateDate = ""
	v.CreateMemo = ""
	v.EndDate = ""
	v.EndReason = ""
	v.Ender = ""
	v.MainOutOrderId = ""
	v.ModifiedMemo = ""
	v.ModifiedReasonText = ""
	v.OutOrderId = ""
	v.ReasonText = ""
	v.RequestId = ""
	v.ReverseStatusStr = ""
	v.ReverseTypeStr = ""
	v.StoreId = ""
	v.BizOrderId = 0
	v.CreateChannel = 0
	v.FetchOrderId = 0
	v.MainBizOrderId = 0
	v.ModifiedReasonId = 0
	v.ReasonId = 0
	v.RefundAmount = 0
	v.ReverseId = 0
	v.ReverseStatus = 0
	v.ReverseType = 0
	v.SellerId = 0
	v.Weight = false
	poolLoadReverseResponse.Put(v)
}
