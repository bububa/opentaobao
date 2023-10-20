package qimen

import (
	"sync"
)

// TaobaoQimenDeliveryorderCreateDetail 结构体
type TaobaoQimenDeliveryorderCreateDetail struct {
	// 商品列表
	Items []Item `json:"items,omitempty" xml:"items>item,omitempty"`
}

var poolTaobaoQimenDeliveryorderCreateDetail = sync.Pool{
	New: func() any {
		return new(TaobaoQimenDeliveryorderCreateDetail)
	},
}

// GetTaobaoQimenDeliveryorderCreateDetail() 从对象池中获取TaobaoQimenDeliveryorderCreateDetail
func GetTaobaoQimenDeliveryorderCreateDetail() *TaobaoQimenDeliveryorderCreateDetail {
	return poolTaobaoQimenDeliveryorderCreateDetail.Get().(*TaobaoQimenDeliveryorderCreateDetail)
}

// ReleaseTaobaoQimenDeliveryorderCreateDetail 释放TaobaoQimenDeliveryorderCreateDetail
func ReleaseTaobaoQimenDeliveryorderCreateDetail(v *TaobaoQimenDeliveryorderCreateDetail) {
	v.Items = v.Items[:0]
	poolTaobaoQimenDeliveryorderCreateDetail.Put(v)
}
