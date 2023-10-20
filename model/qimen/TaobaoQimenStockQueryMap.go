package qimen

import (
	"sync"
)

// TaobaoQimenStockQueryMap 结构体
type TaobaoQimenStockQueryMap struct {
}

var poolTaobaoQimenStockQueryMap = sync.Pool{
	New: func() any {
		return new(TaobaoQimenStockQueryMap)
	},
}

// GetTaobaoQimenStockQueryMap() 从对象池中获取TaobaoQimenStockQueryMap
func GetTaobaoQimenStockQueryMap() *TaobaoQimenStockQueryMap {
	return poolTaobaoQimenStockQueryMap.Get().(*TaobaoQimenStockQueryMap)
}

// ReleaseTaobaoQimenStockQueryMap 释放TaobaoQimenStockQueryMap
func ReleaseTaobaoQimenStockQueryMap(v *TaobaoQimenStockQueryMap) {
	poolTaobaoQimenStockQueryMap.Put(v)
}
