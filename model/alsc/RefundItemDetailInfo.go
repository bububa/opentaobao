package alsc

import (
	"sync"
)

// RefundItemDetailInfo 结构体
type RefundItemDetailInfo struct {
	// 商品id
	OutItemId string `json:"out_item_id,omitempty" xml:"out_item_id,omitempty"`
	// 商品名称
	OutItemName string `json:"out_item_name,omitempty" xml:"out_item_name,omitempty"`
	// 退款单商品明细id
	OutRefundItemDetailNo string `json:"out_refund_item_detail_no,omitempty" xml:"out_refund_item_detail_no,omitempty"`
	// 外部退款单id
	OutRefundOrderNo string `json:"out_refund_order_no,omitempty" xml:"out_refund_order_no,omitempty"`
	// 外部子订单号
	OutSubOrderNo string `json:"out_sub_order_no,omitempty" xml:"out_sub_order_no,omitempty"`
	// 退款商品数
	RefundCount int64 `json:"refund_count,omitempty" xml:"refund_count,omitempty"`
}

var poolRefundItemDetailInfo = sync.Pool{
	New: func() any {
		return new(RefundItemDetailInfo)
	},
}

// GetRefundItemDetailInfo() 从对象池中获取RefundItemDetailInfo
func GetRefundItemDetailInfo() *RefundItemDetailInfo {
	return poolRefundItemDetailInfo.Get().(*RefundItemDetailInfo)
}

// ReleaseRefundItemDetailInfo 释放RefundItemDetailInfo
func ReleaseRefundItemDetailInfo(v *RefundItemDetailInfo) {
	v.OutItemId = ""
	v.OutItemName = ""
	v.OutRefundItemDetailNo = ""
	v.OutRefundOrderNo = ""
	v.OutSubOrderNo = ""
	v.RefundCount = 0
	poolRefundItemDetailInfo.Put(v)
}
