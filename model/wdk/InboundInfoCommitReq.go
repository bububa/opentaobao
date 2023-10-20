package wdk

import (
	"sync"
)

// InboundInfoCommitReq 结构体
type InboundInfoCommitReq struct {
	// 入库商品明细
	InboundItemInfos []InboundItemInfo `json:"inbound_item_infos,omitempty" xml:"inbound_item_infos>inbound_item_info,omitempty"`
	// 收货入库单号
	InboundOrderNo string `json:"inbound_order_no,omitempty" xml:"inbound_order_no,omitempty"`
	// 收货时间
	ReceivedTime string `json:"received_time,omitempty" xml:"received_time,omitempty"`
	// 采购退货单单号
	ReturnOrderNo string `json:"return_order_no,omitempty" xml:"return_order_no,omitempty"`
	// 商家编码
	MerchantCode string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
}

var poolInboundInfoCommitReq = sync.Pool{
	New: func() any {
		return new(InboundInfoCommitReq)
	},
}

// GetInboundInfoCommitReq() 从对象池中获取InboundInfoCommitReq
func GetInboundInfoCommitReq() *InboundInfoCommitReq {
	return poolInboundInfoCommitReq.Get().(*InboundInfoCommitReq)
}

// ReleaseInboundInfoCommitReq 释放InboundInfoCommitReq
func ReleaseInboundInfoCommitReq(v *InboundInfoCommitReq) {
	v.InboundItemInfos = v.InboundItemInfos[:0]
	v.InboundOrderNo = ""
	v.ReceivedTime = ""
	v.ReturnOrderNo = ""
	v.MerchantCode = ""
	poolInboundInfoCommitReq.Put(v)
}
