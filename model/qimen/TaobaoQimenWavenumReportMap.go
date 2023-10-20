package qimen

import (
	"sync"
)

// TaobaoQimenWavenumReportMap 结构体
type TaobaoQimenWavenumReportMap struct {
}

var poolTaobaoQimenWavenumReportMap = sync.Pool{
	New: func() any {
		return new(TaobaoQimenWavenumReportMap)
	},
}

// GetTaobaoQimenWavenumReportMap() 从对象池中获取TaobaoQimenWavenumReportMap
func GetTaobaoQimenWavenumReportMap() *TaobaoQimenWavenumReportMap {
	return poolTaobaoQimenWavenumReportMap.Get().(*TaobaoQimenWavenumReportMap)
}

// ReleaseTaobaoQimenWavenumReportMap 释放TaobaoQimenWavenumReportMap
func ReleaseTaobaoQimenWavenumReportMap(v *TaobaoQimenWavenumReportMap) {
	poolTaobaoQimenWavenumReportMap.Put(v)
}
