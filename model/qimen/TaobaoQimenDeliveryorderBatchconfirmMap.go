package qimen

import (
	"sync"
)

// TaobaoQimenDeliveryorderBatchconfirmMap 结构体
type TaobaoQimenDeliveryorderBatchconfirmMap struct {
}

var poolTaobaoQimenDeliveryorderBatchconfirmMap = sync.Pool{
	New: func() any {
		return new(TaobaoQimenDeliveryorderBatchconfirmMap)
	},
}

// GetTaobaoQimenDeliveryorderBatchconfirmMap() 从对象池中获取TaobaoQimenDeliveryorderBatchconfirmMap
func GetTaobaoQimenDeliveryorderBatchconfirmMap() *TaobaoQimenDeliveryorderBatchconfirmMap {
	return poolTaobaoQimenDeliveryorderBatchconfirmMap.Get().(*TaobaoQimenDeliveryorderBatchconfirmMap)
}

// ReleaseTaobaoQimenDeliveryorderBatchconfirmMap 释放TaobaoQimenDeliveryorderBatchconfirmMap
func ReleaseTaobaoQimenDeliveryorderBatchconfirmMap(v *TaobaoQimenDeliveryorderBatchconfirmMap) {
	poolTaobaoQimenDeliveryorderBatchconfirmMap.Put(v)
}
