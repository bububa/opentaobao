package qimen

import (
	"sync"
)

// TaobaoQimenWarehouseinfoQueryMap 结构体
type TaobaoQimenWarehouseinfoQueryMap struct {
}

var poolTaobaoQimenWarehouseinfoQueryMap = sync.Pool{
	New: func() any {
		return new(TaobaoQimenWarehouseinfoQueryMap)
	},
}

// GetTaobaoQimenWarehouseinfoQueryMap() 从对象池中获取TaobaoQimenWarehouseinfoQueryMap
func GetTaobaoQimenWarehouseinfoQueryMap() *TaobaoQimenWarehouseinfoQueryMap {
	return poolTaobaoQimenWarehouseinfoQueryMap.Get().(*TaobaoQimenWarehouseinfoQueryMap)
}

// ReleaseTaobaoQimenWarehouseinfoQueryMap 释放TaobaoQimenWarehouseinfoQueryMap
func ReleaseTaobaoQimenWarehouseinfoQueryMap(v *TaobaoQimenWarehouseinfoQueryMap) {
	poolTaobaoQimenWarehouseinfoQueryMap.Put(v)
}
