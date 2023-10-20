package qimen

import (
	"sync"
)

// TaobaoQimenInventoryQueryMap 结构体
type TaobaoQimenInventoryQueryMap struct {
}

var poolTaobaoQimenInventoryQueryMap = sync.Pool{
	New: func() any {
		return new(TaobaoQimenInventoryQueryMap)
	},
}

// GetTaobaoQimenInventoryQueryMap() 从对象池中获取TaobaoQimenInventoryQueryMap
func GetTaobaoQimenInventoryQueryMap() *TaobaoQimenInventoryQueryMap {
	return poolTaobaoQimenInventoryQueryMap.Get().(*TaobaoQimenInventoryQueryMap)
}

// ReleaseTaobaoQimenInventoryQueryMap 释放TaobaoQimenInventoryQueryMap
func ReleaseTaobaoQimenInventoryQueryMap(v *TaobaoQimenInventoryQueryMap) {
	poolTaobaoQimenInventoryQueryMap.Put(v)
}
