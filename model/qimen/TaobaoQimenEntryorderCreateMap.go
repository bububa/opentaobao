package qimen

import (
	"sync"
)

// TaobaoQimenEntryorderCreateMap 结构体
type TaobaoQimenEntryorderCreateMap struct {
}

var poolTaobaoQimenEntryorderCreateMap = sync.Pool{
	New: func() any {
		return new(TaobaoQimenEntryorderCreateMap)
	},
}

// GetTaobaoQimenEntryorderCreateMap() 从对象池中获取TaobaoQimenEntryorderCreateMap
func GetTaobaoQimenEntryorderCreateMap() *TaobaoQimenEntryorderCreateMap {
	return poolTaobaoQimenEntryorderCreateMap.Get().(*TaobaoQimenEntryorderCreateMap)
}

// ReleaseTaobaoQimenEntryorderCreateMap 释放TaobaoQimenEntryorderCreateMap
func ReleaseTaobaoQimenEntryorderCreateMap(v *TaobaoQimenEntryorderCreateMap) {
	poolTaobaoQimenEntryorderCreateMap.Put(v)
}
