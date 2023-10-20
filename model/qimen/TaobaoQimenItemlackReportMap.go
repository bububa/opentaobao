package qimen

import (
	"sync"
)

// TaobaoQimenItemlackReportMap 结构体
type TaobaoQimenItemlackReportMap struct {
}

var poolTaobaoQimenItemlackReportMap = sync.Pool{
	New: func() any {
		return new(TaobaoQimenItemlackReportMap)
	},
}

// GetTaobaoQimenItemlackReportMap() 从对象池中获取TaobaoQimenItemlackReportMap
func GetTaobaoQimenItemlackReportMap() *TaobaoQimenItemlackReportMap {
	return poolTaobaoQimenItemlackReportMap.Get().(*TaobaoQimenItemlackReportMap)
}

// ReleaseTaobaoQimenItemlackReportMap 释放TaobaoQimenItemlackReportMap
func ReleaseTaobaoQimenItemlackReportMap(v *TaobaoQimenItemlackReportMap) {
	poolTaobaoQimenItemlackReportMap.Put(v)
}
