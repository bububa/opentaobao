package qimen

import (
	"sync"
)

// TaobaoQimenOrderprocessQueryMap 结构体
type TaobaoQimenOrderprocessQueryMap struct {
}

var poolTaobaoQimenOrderprocessQueryMap = sync.Pool{
	New: func() any {
		return new(TaobaoQimenOrderprocessQueryMap)
	},
}

// GetTaobaoQimenOrderprocessQueryMap() 从对象池中获取TaobaoQimenOrderprocessQueryMap
func GetTaobaoQimenOrderprocessQueryMap() *TaobaoQimenOrderprocessQueryMap {
	return poolTaobaoQimenOrderprocessQueryMap.Get().(*TaobaoQimenOrderprocessQueryMap)
}

// ReleaseTaobaoQimenOrderprocessQueryMap 释放TaobaoQimenOrderprocessQueryMap
func ReleaseTaobaoQimenOrderprocessQueryMap(v *TaobaoQimenOrderprocessQueryMap) {
	poolTaobaoQimenOrderprocessQueryMap.Put(v)
}
