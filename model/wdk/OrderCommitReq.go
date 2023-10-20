package wdk

import (
	"sync"
)

// OrderCommitReq 结构体
type OrderCommitReq struct {
	// 商品列表
	ItemConfirmInfos []ItemConfirmInfo `json:"item_confirm_infos,omitempty" xml:"item_confirm_infos>item_confirm_info,omitempty"`
	// 订单号
	ExternalOrderNo string `json:"external_order_no,omitempty" xml:"external_order_no,omitempty"`
	// 商家编码
	MerchantCode string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
}

var poolOrderCommitReq = sync.Pool{
	New: func() any {
		return new(OrderCommitReq)
	},
}

// GetOrderCommitReq() 从对象池中获取OrderCommitReq
func GetOrderCommitReq() *OrderCommitReq {
	return poolOrderCommitReq.Get().(*OrderCommitReq)
}

// ReleaseOrderCommitReq 释放OrderCommitReq
func ReleaseOrderCommitReq(v *OrderCommitReq) {
	v.ItemConfirmInfos = v.ItemConfirmInfos[:0]
	v.ExternalOrderNo = ""
	v.MerchantCode = ""
	poolOrderCommitReq.Put(v)
}
