package qimen

import (
	"sync"
)

// TaobaoQimenDeliveryorderBatchcreateMap 结构体
type TaobaoQimenDeliveryorderBatchcreateMap struct {
}

var poolTaobaoQimenDeliveryorderBatchcreateMap = sync.Pool{
	New: func() any {
		return new(TaobaoQimenDeliveryorderBatchcreateMap)
	},
}

// GetTaobaoQimenDeliveryorderBatchcreateMap() 从对象池中获取TaobaoQimenDeliveryorderBatchcreateMap
func GetTaobaoQimenDeliveryorderBatchcreateMap() *TaobaoQimenDeliveryorderBatchcreateMap {
	return poolTaobaoQimenDeliveryorderBatchcreateMap.Get().(*TaobaoQimenDeliveryorderBatchcreateMap)
}

// ReleaseTaobaoQimenDeliveryorderBatchcreateMap 释放TaobaoQimenDeliveryorderBatchcreateMap
func ReleaseTaobaoQimenDeliveryorderBatchcreateMap(v *TaobaoQimenDeliveryorderBatchcreateMap) {
	poolTaobaoQimenDeliveryorderBatchcreateMap.Put(v)
}
