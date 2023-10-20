package qimen

import (
	"sync"
)

// TaobaoQimenEntryorderConfirmMap 结构体
type TaobaoQimenEntryorderConfirmMap struct {
}

var poolTaobaoQimenEntryorderConfirmMap = sync.Pool{
	New: func() any {
		return new(TaobaoQimenEntryorderConfirmMap)
	},
}

// GetTaobaoQimenEntryorderConfirmMap() 从对象池中获取TaobaoQimenEntryorderConfirmMap
func GetTaobaoQimenEntryorderConfirmMap() *TaobaoQimenEntryorderConfirmMap {
	return poolTaobaoQimenEntryorderConfirmMap.Get().(*TaobaoQimenEntryorderConfirmMap)
}

// ReleaseTaobaoQimenEntryorderConfirmMap 释放TaobaoQimenEntryorderConfirmMap
func ReleaseTaobaoQimenEntryorderConfirmMap(v *TaobaoQimenEntryorderConfirmMap) {
	poolTaobaoQimenEntryorderConfirmMap.Put(v)
}
