package qimen

import (
	"sync"
)

// TaobaoQimenOrderPendingMap 结构体
type TaobaoQimenOrderPendingMap struct {
}

var poolTaobaoQimenOrderPendingMap = sync.Pool{
	New: func() any {
		return new(TaobaoQimenOrderPendingMap)
	},
}

// GetTaobaoQimenOrderPendingMap() 从对象池中获取TaobaoQimenOrderPendingMap
func GetTaobaoQimenOrderPendingMap() *TaobaoQimenOrderPendingMap {
	return poolTaobaoQimenOrderPendingMap.Get().(*TaobaoQimenOrderPendingMap)
}

// ReleaseTaobaoQimenOrderPendingMap 释放TaobaoQimenOrderPendingMap
func ReleaseTaobaoQimenOrderPendingMap(v *TaobaoQimenOrderPendingMap) {
	poolTaobaoQimenOrderPendingMap.Put(v)
}
