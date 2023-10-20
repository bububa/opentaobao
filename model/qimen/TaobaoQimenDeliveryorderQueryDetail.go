package qimen

import (
	"sync"
)

// TaobaoQimenDeliveryorderQueryDetail 结构体
type TaobaoQimenDeliveryorderQueryDetail struct {
	// 商品列表
	Items []Item `json:"items,omitempty" xml:"items>item,omitempty"`
}

var poolTaobaoQimenDeliveryorderQueryDetail = sync.Pool{
	New: func() any {
		return new(TaobaoQimenDeliveryorderQueryDetail)
	},
}

// GetTaobaoQimenDeliveryorderQueryDetail() 从对象池中获取TaobaoQimenDeliveryorderQueryDetail
func GetTaobaoQimenDeliveryorderQueryDetail() *TaobaoQimenDeliveryorderQueryDetail {
	return poolTaobaoQimenDeliveryorderQueryDetail.Get().(*TaobaoQimenDeliveryorderQueryDetail)
}

// ReleaseTaobaoQimenDeliveryorderQueryDetail 释放TaobaoQimenDeliveryorderQueryDetail
func ReleaseTaobaoQimenDeliveryorderQueryDetail(v *TaobaoQimenDeliveryorderQueryDetail) {
	v.Items = v.Items[:0]
	poolTaobaoQimenDeliveryorderQueryDetail.Put(v)
}
