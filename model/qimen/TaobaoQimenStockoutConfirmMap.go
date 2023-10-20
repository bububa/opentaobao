package qimen

import (
	"sync"
)

// TaobaoQimenStockoutConfirmMap 结构体
type TaobaoQimenStockoutConfirmMap struct {
}

var poolTaobaoQimenStockoutConfirmMap = sync.Pool{
	New: func() any {
		return new(TaobaoQimenStockoutConfirmMap)
	},
}

// GetTaobaoQimenStockoutConfirmMap() 从对象池中获取TaobaoQimenStockoutConfirmMap
func GetTaobaoQimenStockoutConfirmMap() *TaobaoQimenStockoutConfirmMap {
	return poolTaobaoQimenStockoutConfirmMap.Get().(*TaobaoQimenStockoutConfirmMap)
}

// ReleaseTaobaoQimenStockoutConfirmMap 释放TaobaoQimenStockoutConfirmMap
func ReleaseTaobaoQimenStockoutConfirmMap(v *TaobaoQimenStockoutConfirmMap) {
	poolTaobaoQimenStockoutConfirmMap.Put(v)
}
