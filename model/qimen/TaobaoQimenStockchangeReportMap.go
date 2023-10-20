package qimen

import (
	"sync"
)

// TaobaoQimenStockchangeReportMap 结构体
type TaobaoQimenStockchangeReportMap struct {
}

var poolTaobaoQimenStockchangeReportMap = sync.Pool{
	New: func() any {
		return new(TaobaoQimenStockchangeReportMap)
	},
}

// GetTaobaoQimenStockchangeReportMap() 从对象池中获取TaobaoQimenStockchangeReportMap
func GetTaobaoQimenStockchangeReportMap() *TaobaoQimenStockchangeReportMap {
	return poolTaobaoQimenStockchangeReportMap.Get().(*TaobaoQimenStockchangeReportMap)
}

// ReleaseTaobaoQimenStockchangeReportMap 释放TaobaoQimenStockchangeReportMap
func ReleaseTaobaoQimenStockchangeReportMap(v *TaobaoQimenStockchangeReportMap) {
	poolTaobaoQimenStockchangeReportMap.Put(v)
}
