package qimen

import (
	"sync"
)

// TaobaoQimenDeliveryorderBatchconfirmDetail 结构体
type TaobaoQimenDeliveryorderBatchconfirmDetail struct {
	// 订单商品列表
	Items []Item `json:"items,omitempty" xml:"items>item,omitempty"`
}

var poolTaobaoQimenDeliveryorderBatchconfirmDetail = sync.Pool{
	New: func() any {
		return new(TaobaoQimenDeliveryorderBatchconfirmDetail)
	},
}

// GetTaobaoQimenDeliveryorderBatchconfirmDetail() 从对象池中获取TaobaoQimenDeliveryorderBatchconfirmDetail
func GetTaobaoQimenDeliveryorderBatchconfirmDetail() *TaobaoQimenDeliveryorderBatchconfirmDetail {
	return poolTaobaoQimenDeliveryorderBatchconfirmDetail.Get().(*TaobaoQimenDeliveryorderBatchconfirmDetail)
}

// ReleaseTaobaoQimenDeliveryorderBatchconfirmDetail 释放TaobaoQimenDeliveryorderBatchconfirmDetail
func ReleaseTaobaoQimenDeliveryorderBatchconfirmDetail(v *TaobaoQimenDeliveryorderBatchconfirmDetail) {
	v.Items = v.Items[:0]
	poolTaobaoQimenDeliveryorderBatchconfirmDetail.Put(v)
}
