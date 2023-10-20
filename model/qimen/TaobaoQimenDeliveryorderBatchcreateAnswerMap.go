package qimen

import (
	"sync"
)

// TaobaoQimenDeliveryorderBatchcreateAnswerMap 结构体
type TaobaoQimenDeliveryorderBatchcreateAnswerMap struct {
}

var poolTaobaoQimenDeliveryorderBatchcreateAnswerMap = sync.Pool{
	New: func() any {
		return new(TaobaoQimenDeliveryorderBatchcreateAnswerMap)
	},
}

// GetTaobaoQimenDeliveryorderBatchcreateAnswerMap() 从对象池中获取TaobaoQimenDeliveryorderBatchcreateAnswerMap
func GetTaobaoQimenDeliveryorderBatchcreateAnswerMap() *TaobaoQimenDeliveryorderBatchcreateAnswerMap {
	return poolTaobaoQimenDeliveryorderBatchcreateAnswerMap.Get().(*TaobaoQimenDeliveryorderBatchcreateAnswerMap)
}

// ReleaseTaobaoQimenDeliveryorderBatchcreateAnswerMap 释放TaobaoQimenDeliveryorderBatchcreateAnswerMap
func ReleaseTaobaoQimenDeliveryorderBatchcreateAnswerMap(v *TaobaoQimenDeliveryorderBatchcreateAnswerMap) {
	poolTaobaoQimenDeliveryorderBatchcreateAnswerMap.Put(v)
}
