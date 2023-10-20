package qimen

import (
	"sync"
)

// TaobaoQimenCombineitemSynchronizeMap 结构体
type TaobaoQimenCombineitemSynchronizeMap struct {
}

var poolTaobaoQimenCombineitemSynchronizeMap = sync.Pool{
	New: func() any {
		return new(TaobaoQimenCombineitemSynchronizeMap)
	},
}

// GetTaobaoQimenCombineitemSynchronizeMap() 从对象池中获取TaobaoQimenCombineitemSynchronizeMap
func GetTaobaoQimenCombineitemSynchronizeMap() *TaobaoQimenCombineitemSynchronizeMap {
	return poolTaobaoQimenCombineitemSynchronizeMap.Get().(*TaobaoQimenCombineitemSynchronizeMap)
}

// ReleaseTaobaoQimenCombineitemSynchronizeMap 释放TaobaoQimenCombineitemSynchronizeMap
func ReleaseTaobaoQimenCombineitemSynchronizeMap(v *TaobaoQimenCombineitemSynchronizeMap) {
	poolTaobaoQimenCombineitemSynchronizeMap.Put(v)
}
