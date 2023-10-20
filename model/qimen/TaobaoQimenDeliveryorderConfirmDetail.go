package qimen

import (
	"sync"
)

// TaobaoQimenDeliveryorderConfirmDetail 结构体
type TaobaoQimenDeliveryorderConfirmDetail struct {
	// 商品列表
	Items []Item `json:"items,omitempty" xml:"items>item,omitempty"`
}

var poolTaobaoQimenDeliveryorderConfirmDetail = sync.Pool{
	New: func() any {
		return new(TaobaoQimenDeliveryorderConfirmDetail)
	},
}

// GetTaobaoQimenDeliveryorderConfirmDetail() 从对象池中获取TaobaoQimenDeliveryorderConfirmDetail
func GetTaobaoQimenDeliveryorderConfirmDetail() *TaobaoQimenDeliveryorderConfirmDetail {
	return poolTaobaoQimenDeliveryorderConfirmDetail.Get().(*TaobaoQimenDeliveryorderConfirmDetail)
}

// ReleaseTaobaoQimenDeliveryorderConfirmDetail 释放TaobaoQimenDeliveryorderConfirmDetail
func ReleaseTaobaoQimenDeliveryorderConfirmDetail(v *TaobaoQimenDeliveryorderConfirmDetail) {
	v.Items = v.Items[:0]
	poolTaobaoQimenDeliveryorderConfirmDetail.Put(v)
}
