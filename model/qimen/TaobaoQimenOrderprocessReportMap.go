package qimen

import (
	"sync"
)

// TaobaoQimenOrderprocessReportMap 结构体
type TaobaoQimenOrderprocessReportMap struct {
}

var poolTaobaoQimenOrderprocessReportMap = sync.Pool{
	New: func() any {
		return new(TaobaoQimenOrderprocessReportMap)
	},
}

// GetTaobaoQimenOrderprocessReportMap() 从对象池中获取TaobaoQimenOrderprocessReportMap
func GetTaobaoQimenOrderprocessReportMap() *TaobaoQimenOrderprocessReportMap {
	return poolTaobaoQimenOrderprocessReportMap.Get().(*TaobaoQimenOrderprocessReportMap)
}

// ReleaseTaobaoQimenOrderprocessReportMap 释放TaobaoQimenOrderprocessReportMap
func ReleaseTaobaoQimenOrderprocessReportMap(v *TaobaoQimenOrderprocessReportMap) {
	poolTaobaoQimenOrderprocessReportMap.Put(v)
}
