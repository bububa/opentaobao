package qimen

import (
	"sync"
)

// TaobaoQimenDeliveryorderConfirmMap 结构体
type TaobaoQimenDeliveryorderConfirmMap struct {
}

var poolTaobaoQimenDeliveryorderConfirmMap = sync.Pool{
	New: func() any {
		return new(TaobaoQimenDeliveryorderConfirmMap)
	},
}

// GetTaobaoQimenDeliveryorderConfirmMap() 从对象池中获取TaobaoQimenDeliveryorderConfirmMap
func GetTaobaoQimenDeliveryorderConfirmMap() *TaobaoQimenDeliveryorderConfirmMap {
	return poolTaobaoQimenDeliveryorderConfirmMap.Get().(*TaobaoQimenDeliveryorderConfirmMap)
}

// ReleaseTaobaoQimenDeliveryorderConfirmMap 释放TaobaoQimenDeliveryorderConfirmMap
func ReleaseTaobaoQimenDeliveryorderConfirmMap(v *TaobaoQimenDeliveryorderConfirmMap) {
	poolTaobaoQimenDeliveryorderConfirmMap.Put(v)
}
