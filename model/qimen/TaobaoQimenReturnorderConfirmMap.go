package qimen

import (
	"sync"
)

// TaobaoQimenReturnorderConfirmMap 结构体
type TaobaoQimenReturnorderConfirmMap struct {
}

var poolTaobaoQimenReturnorderConfirmMap = sync.Pool{
	New: func() any {
		return new(TaobaoQimenReturnorderConfirmMap)
	},
}

// GetTaobaoQimenReturnorderConfirmMap() 从对象池中获取TaobaoQimenReturnorderConfirmMap
func GetTaobaoQimenReturnorderConfirmMap() *TaobaoQimenReturnorderConfirmMap {
	return poolTaobaoQimenReturnorderConfirmMap.Get().(*TaobaoQimenReturnorderConfirmMap)
}

// ReleaseTaobaoQimenReturnorderConfirmMap 释放TaobaoQimenReturnorderConfirmMap
func ReleaseTaobaoQimenReturnorderConfirmMap(v *TaobaoQimenReturnorderConfirmMap) {
	poolTaobaoQimenReturnorderConfirmMap.Put(v)
}
