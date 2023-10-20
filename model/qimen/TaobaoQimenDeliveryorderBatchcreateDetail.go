package qimen

import (
	"sync"
)

// TaobaoQimenDeliveryorderBatchcreateDetail 结构体
type TaobaoQimenDeliveryorderBatchcreateDetail struct {
	// 商品列表
	Items []Item `json:"items,omitempty" xml:"items>item,omitempty"`
}

var poolTaobaoQimenDeliveryorderBatchcreateDetail = sync.Pool{
	New: func() any {
		return new(TaobaoQimenDeliveryorderBatchcreateDetail)
	},
}

// GetTaobaoQimenDeliveryorderBatchcreateDetail() 从对象池中获取TaobaoQimenDeliveryorderBatchcreateDetail
func GetTaobaoQimenDeliveryorderBatchcreateDetail() *TaobaoQimenDeliveryorderBatchcreateDetail {
	return poolTaobaoQimenDeliveryorderBatchcreateDetail.Get().(*TaobaoQimenDeliveryorderBatchcreateDetail)
}

// ReleaseTaobaoQimenDeliveryorderBatchcreateDetail 释放TaobaoQimenDeliveryorderBatchcreateDetail
func ReleaseTaobaoQimenDeliveryorderBatchcreateDetail(v *TaobaoQimenDeliveryorderBatchcreateDetail) {
	v.Items = v.Items[:0]
	poolTaobaoQimenDeliveryorderBatchcreateDetail.Put(v)
}
