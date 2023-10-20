package qimen

import (
	"sync"
)

// TaobaoQimenInventoryReportMap 结构体
type TaobaoQimenInventoryReportMap struct {
}

var poolTaobaoQimenInventoryReportMap = sync.Pool{
	New: func() any {
		return new(TaobaoQimenInventoryReportMap)
	},
}

// GetTaobaoQimenInventoryReportMap() 从对象池中获取TaobaoQimenInventoryReportMap
func GetTaobaoQimenInventoryReportMap() *TaobaoQimenInventoryReportMap {
	return poolTaobaoQimenInventoryReportMap.Get().(*TaobaoQimenInventoryReportMap)
}

// ReleaseTaobaoQimenInventoryReportMap 释放TaobaoQimenInventoryReportMap
func ReleaseTaobaoQimenInventoryReportMap(v *TaobaoQimenInventoryReportMap) {
	poolTaobaoQimenInventoryReportMap.Put(v)
}
