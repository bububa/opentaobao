package qimen

import (
	"sync"
)

// TaobaoQimenDeliveryorderCreateMap 结构体
type TaobaoQimenDeliveryorderCreateMap struct {
}

var poolTaobaoQimenDeliveryorderCreateMap = sync.Pool{
	New: func() any {
		return new(TaobaoQimenDeliveryorderCreateMap)
	},
}

// GetTaobaoQimenDeliveryorderCreateMap() 从对象池中获取TaobaoQimenDeliveryorderCreateMap
func GetTaobaoQimenDeliveryorderCreateMap() *TaobaoQimenDeliveryorderCreateMap {
	return poolTaobaoQimenDeliveryorderCreateMap.Get().(*TaobaoQimenDeliveryorderCreateMap)
}

// ReleaseTaobaoQimenDeliveryorderCreateMap 释放TaobaoQimenDeliveryorderCreateMap
func ReleaseTaobaoQimenDeliveryorderCreateMap(v *TaobaoQimenDeliveryorderCreateMap) {
	poolTaobaoQimenDeliveryorderCreateMap.Put(v)
}
