package qimen

import (
	"sync"
)

// TaobaoQimenEntryorderQueryMap 结构体
type TaobaoQimenEntryorderQueryMap struct {
}

var poolTaobaoQimenEntryorderQueryMap = sync.Pool{
	New: func() any {
		return new(TaobaoQimenEntryorderQueryMap)
	},
}

// GetTaobaoQimenEntryorderQueryMap() 从对象池中获取TaobaoQimenEntryorderQueryMap
func GetTaobaoQimenEntryorderQueryMap() *TaobaoQimenEntryorderQueryMap {
	return poolTaobaoQimenEntryorderQueryMap.Get().(*TaobaoQimenEntryorderQueryMap)
}

// ReleaseTaobaoQimenEntryorderQueryMap 释放TaobaoQimenEntryorderQueryMap
func ReleaseTaobaoQimenEntryorderQueryMap(v *TaobaoQimenEntryorderQueryMap) {
	poolTaobaoQimenEntryorderQueryMap.Put(v)
}
