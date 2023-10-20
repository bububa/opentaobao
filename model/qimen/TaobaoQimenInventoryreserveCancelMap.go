package qimen

import (
	"sync"
)

// TaobaoQimenInventoryreserveCancelMap 结构体
type TaobaoQimenInventoryreserveCancelMap struct {
}

var poolTaobaoQimenInventoryreserveCancelMap = sync.Pool{
	New: func() any {
		return new(TaobaoQimenInventoryreserveCancelMap)
	},
}

// GetTaobaoQimenInventoryreserveCancelMap() 从对象池中获取TaobaoQimenInventoryreserveCancelMap
func GetTaobaoQimenInventoryreserveCancelMap() *TaobaoQimenInventoryreserveCancelMap {
	return poolTaobaoQimenInventoryreserveCancelMap.Get().(*TaobaoQimenInventoryreserveCancelMap)
}

// ReleaseTaobaoQimenInventoryreserveCancelMap 释放TaobaoQimenInventoryreserveCancelMap
func ReleaseTaobaoQimenInventoryreserveCancelMap(v *TaobaoQimenInventoryreserveCancelMap) {
	poolTaobaoQimenInventoryreserveCancelMap.Put(v)
}
