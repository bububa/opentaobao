package qimen

import (
	"sync"
)

// TaobaoQimenItemsSynchronizeMap 结构体
type TaobaoQimenItemsSynchronizeMap struct {
}

var poolTaobaoQimenItemsSynchronizeMap = sync.Pool{
	New: func() any {
		return new(TaobaoQimenItemsSynchronizeMap)
	},
}

// GetTaobaoQimenItemsSynchronizeMap() 从对象池中获取TaobaoQimenItemsSynchronizeMap
func GetTaobaoQimenItemsSynchronizeMap() *TaobaoQimenItemsSynchronizeMap {
	return poolTaobaoQimenItemsSynchronizeMap.Get().(*TaobaoQimenItemsSynchronizeMap)
}

// ReleaseTaobaoQimenItemsSynchronizeMap 释放TaobaoQimenItemsSynchronizeMap
func ReleaseTaobaoQimenItemsSynchronizeMap(v *TaobaoQimenItemsSynchronizeMap) {
	poolTaobaoQimenItemsSynchronizeMap.Put(v)
}
