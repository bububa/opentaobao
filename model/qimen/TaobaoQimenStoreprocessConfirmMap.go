package qimen

import (
	"sync"
)

// TaobaoQimenStoreprocessConfirmMap 结构体
type TaobaoQimenStoreprocessConfirmMap struct {
}

var poolTaobaoQimenStoreprocessConfirmMap = sync.Pool{
	New: func() any {
		return new(TaobaoQimenStoreprocessConfirmMap)
	},
}

// GetTaobaoQimenStoreprocessConfirmMap() 从对象池中获取TaobaoQimenStoreprocessConfirmMap
func GetTaobaoQimenStoreprocessConfirmMap() *TaobaoQimenStoreprocessConfirmMap {
	return poolTaobaoQimenStoreprocessConfirmMap.Get().(*TaobaoQimenStoreprocessConfirmMap)
}

// ReleaseTaobaoQimenStoreprocessConfirmMap 释放TaobaoQimenStoreprocessConfirmMap
func ReleaseTaobaoQimenStoreprocessConfirmMap(v *TaobaoQimenStoreprocessConfirmMap) {
	poolTaobaoQimenStoreprocessConfirmMap.Put(v)
}
