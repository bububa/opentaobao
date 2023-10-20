package qimen

import (
	"sync"
)

// TaobaoQimenStoreprocessCreateMap 结构体
type TaobaoQimenStoreprocessCreateMap struct {
}

var poolTaobaoQimenStoreprocessCreateMap = sync.Pool{
	New: func() any {
		return new(TaobaoQimenStoreprocessCreateMap)
	},
}

// GetTaobaoQimenStoreprocessCreateMap() 从对象池中获取TaobaoQimenStoreprocessCreateMap
func GetTaobaoQimenStoreprocessCreateMap() *TaobaoQimenStoreprocessCreateMap {
	return poolTaobaoQimenStoreprocessCreateMap.Get().(*TaobaoQimenStoreprocessCreateMap)
}

// ReleaseTaobaoQimenStoreprocessCreateMap 释放TaobaoQimenStoreprocessCreateMap
func ReleaseTaobaoQimenStoreprocessCreateMap(v *TaobaoQimenStoreprocessCreateMap) {
	poolTaobaoQimenStoreprocessCreateMap.Put(v)
}
