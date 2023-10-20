package qimen

import (
	"sync"
)

// TaobaoQimenOrderCancelMap 结构体
type TaobaoQimenOrderCancelMap struct {
}

var poolTaobaoQimenOrderCancelMap = sync.Pool{
	New: func() any {
		return new(TaobaoQimenOrderCancelMap)
	},
}

// GetTaobaoQimenOrderCancelMap() 从对象池中获取TaobaoQimenOrderCancelMap
func GetTaobaoQimenOrderCancelMap() *TaobaoQimenOrderCancelMap {
	return poolTaobaoQimenOrderCancelMap.Get().(*TaobaoQimenOrderCancelMap)
}

// ReleaseTaobaoQimenOrderCancelMap 释放TaobaoQimenOrderCancelMap
func ReleaseTaobaoQimenOrderCancelMap(v *TaobaoQimenOrderCancelMap) {
	poolTaobaoQimenOrderCancelMap.Put(v)
}
