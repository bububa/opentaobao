package wdk

import (
	"sync"
)

// ApplyReverseReq 结构体
type ApplyReverseReq struct {
	// wdk子单list
	BizOrderIds []int64 `json:"biz_order_ids,omitempty" xml:"biz_order_ids>int64,omitempty"`
	// 礼品卡号
	GiftCardNos []string `json:"gift_card_nos,omitempty" xml:"gift_card_nos>string,omitempty"`
	// 外部子单list(biz_order_ids与sub_out_order_ids 二选一)
	SubOutOrderIds []string `json:"sub_out_order_ids,omitempty" xml:"sub_out_order_ids>string,omitempty"`
	// 门店id
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 操作者
	Operator *OperatorVo `json:"operator,omitempty" xml:"operator,omitempty"`
}

var poolApplyReverseReq = sync.Pool{
	New: func() any {
		return new(ApplyReverseReq)
	},
}

// GetApplyReverseReq() 从对象池中获取ApplyReverseReq
func GetApplyReverseReq() *ApplyReverseReq {
	return poolApplyReverseReq.Get().(*ApplyReverseReq)
}

// ReleaseApplyReverseReq 释放ApplyReverseReq
func ReleaseApplyReverseReq(v *ApplyReverseReq) {
	v.BizOrderIds = v.BizOrderIds[:0]
	v.GiftCardNos = v.GiftCardNos[:0]
	v.SubOutOrderIds = v.SubOutOrderIds[:0]
	v.StoreId = ""
	v.Operator = nil
	poolApplyReverseReq.Put(v)
}
