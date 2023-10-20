package qimen

import (
	"sync"
)

// TaobaoQimenSingleitemSynchronizeMap 结构体
type TaobaoQimenSingleitemSynchronizeMap struct {
}

var poolTaobaoQimenSingleitemSynchronizeMap = sync.Pool{
	New: func() any {
		return new(TaobaoQimenSingleitemSynchronizeMap)
	},
}

// GetTaobaoQimenSingleitemSynchronizeMap() 从对象池中获取TaobaoQimenSingleitemSynchronizeMap
func GetTaobaoQimenSingleitemSynchronizeMap() *TaobaoQimenSingleitemSynchronizeMap {
	return poolTaobaoQimenSingleitemSynchronizeMap.Get().(*TaobaoQimenSingleitemSynchronizeMap)
}

// ReleaseTaobaoQimenSingleitemSynchronizeMap 释放TaobaoQimenSingleitemSynchronizeMap
func ReleaseTaobaoQimenSingleitemSynchronizeMap(v *TaobaoQimenSingleitemSynchronizeMap) {
	poolTaobaoQimenSingleitemSynchronizeMap.Put(v)
}
