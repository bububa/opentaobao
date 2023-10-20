package qimen

import (
	"sync"
)

// TaobaoQimenStockoutCreateMap 结构体
type TaobaoQimenStockoutCreateMap struct {
}

var poolTaobaoQimenStockoutCreateMap = sync.Pool{
	New: func() any {
		return new(TaobaoQimenStockoutCreateMap)
	},
}

// GetTaobaoQimenStockoutCreateMap() 从对象池中获取TaobaoQimenStockoutCreateMap
func GetTaobaoQimenStockoutCreateMap() *TaobaoQimenStockoutCreateMap {
	return poolTaobaoQimenStockoutCreateMap.Get().(*TaobaoQimenStockoutCreateMap)
}

// ReleaseTaobaoQimenStockoutCreateMap 释放TaobaoQimenStockoutCreateMap
func ReleaseTaobaoQimenStockoutCreateMap(v *TaobaoQimenStockoutCreateMap) {
	poolTaobaoQimenStockoutCreateMap.Put(v)
}
