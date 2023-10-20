package qimen

import (
	"sync"
)

// TaobaoQimenReturnorderCreateMap 结构体
type TaobaoQimenReturnorderCreateMap struct {
}

var poolTaobaoQimenReturnorderCreateMap = sync.Pool{
	New: func() any {
		return new(TaobaoQimenReturnorderCreateMap)
	},
}

// GetTaobaoQimenReturnorderCreateMap() 从对象池中获取TaobaoQimenReturnorderCreateMap
func GetTaobaoQimenReturnorderCreateMap() *TaobaoQimenReturnorderCreateMap {
	return poolTaobaoQimenReturnorderCreateMap.Get().(*TaobaoQimenReturnorderCreateMap)
}

// ReleaseTaobaoQimenReturnorderCreateMap 释放TaobaoQimenReturnorderCreateMap
func ReleaseTaobaoQimenReturnorderCreateMap(v *TaobaoQimenReturnorderCreateMap) {
	poolTaobaoQimenReturnorderCreateMap.Put(v)
}
