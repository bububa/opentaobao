package qimen

import (
	"sync"
)

// TaobaoQimenOrderexceptionReportMap 结构体
type TaobaoQimenOrderexceptionReportMap struct {
}

var poolTaobaoQimenOrderexceptionReportMap = sync.Pool{
	New: func() any {
		return new(TaobaoQimenOrderexceptionReportMap)
	},
}

// GetTaobaoQimenOrderexceptionReportMap() 从对象池中获取TaobaoQimenOrderexceptionReportMap
func GetTaobaoQimenOrderexceptionReportMap() *TaobaoQimenOrderexceptionReportMap {
	return poolTaobaoQimenOrderexceptionReportMap.Get().(*TaobaoQimenOrderexceptionReportMap)
}

// ReleaseTaobaoQimenOrderexceptionReportMap 释放TaobaoQimenOrderexceptionReportMap
func ReleaseTaobaoQimenOrderexceptionReportMap(v *TaobaoQimenOrderexceptionReportMap) {
	poolTaobaoQimenOrderexceptionReportMap.Put(v)
}
