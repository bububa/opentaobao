package qimen

import (
	"sync"
)

// TaobaoQimenDeliveryorderQueryMap 结构体
type TaobaoQimenDeliveryorderQueryMap struct {
}

var poolTaobaoQimenDeliveryorderQueryMap = sync.Pool{
	New: func() any {
		return new(TaobaoQimenDeliveryorderQueryMap)
	},
}

// GetTaobaoQimenDeliveryorderQueryMap() 从对象池中获取TaobaoQimenDeliveryorderQueryMap
func GetTaobaoQimenDeliveryorderQueryMap() *TaobaoQimenDeliveryorderQueryMap {
	return poolTaobaoQimenDeliveryorderQueryMap.Get().(*TaobaoQimenDeliveryorderQueryMap)
}

// ReleaseTaobaoQimenDeliveryorderQueryMap 释放TaobaoQimenDeliveryorderQueryMap
func ReleaseTaobaoQimenDeliveryorderQueryMap(v *TaobaoQimenDeliveryorderQueryMap) {
	poolTaobaoQimenDeliveryorderQueryMap.Put(v)
}
