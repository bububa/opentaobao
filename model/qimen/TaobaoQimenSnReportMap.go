package qimen

import (
	"sync"
)

// TaobaoQimenSnReportMap 结构体
type TaobaoQimenSnReportMap struct {
}

var poolTaobaoQimenSnReportMap = sync.Pool{
	New: func() any {
		return new(TaobaoQimenSnReportMap)
	},
}

// GetTaobaoQimenSnReportMap() 从对象池中获取TaobaoQimenSnReportMap
func GetTaobaoQimenSnReportMap() *TaobaoQimenSnReportMap {
	return poolTaobaoQimenSnReportMap.Get().(*TaobaoQimenSnReportMap)
}

// ReleaseTaobaoQimenSnReportMap 释放TaobaoQimenSnReportMap
func ReleaseTaobaoQimenSnReportMap(v *TaobaoQimenSnReportMap) {
	poolTaobaoQimenSnReportMap.Put(v)
}
