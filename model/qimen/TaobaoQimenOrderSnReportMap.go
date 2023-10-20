package qimen

import (
	"sync"
)

// TaobaoQimenOrderSnReportMap 结构体
type TaobaoQimenOrderSnReportMap struct {
}

var poolTaobaoQimenOrderSnReportMap = sync.Pool{
	New: func() any {
		return new(TaobaoQimenOrderSnReportMap)
	},
}

// GetTaobaoQimenOrderSnReportMap() 从对象池中获取TaobaoQimenOrderSnReportMap
func GetTaobaoQimenOrderSnReportMap() *TaobaoQimenOrderSnReportMap {
	return poolTaobaoQimenOrderSnReportMap.Get().(*TaobaoQimenOrderSnReportMap)
}

// ReleaseTaobaoQimenOrderSnReportMap 释放TaobaoQimenOrderSnReportMap
func ReleaseTaobaoQimenOrderSnReportMap(v *TaobaoQimenOrderSnReportMap) {
	poolTaobaoQimenOrderSnReportMap.Put(v)
}
