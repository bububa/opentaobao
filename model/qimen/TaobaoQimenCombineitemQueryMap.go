package qimen

import (
	"sync"
)

// TaobaoQimenCombineitemQueryMap 结构体
type TaobaoQimenCombineitemQueryMap struct {
}

var poolTaobaoQimenCombineitemQueryMap = sync.Pool{
	New: func() any {
		return new(TaobaoQimenCombineitemQueryMap)
	},
}

// GetTaobaoQimenCombineitemQueryMap() 从对象池中获取TaobaoQimenCombineitemQueryMap
func GetTaobaoQimenCombineitemQueryMap() *TaobaoQimenCombineitemQueryMap {
	return poolTaobaoQimenCombineitemQueryMap.Get().(*TaobaoQimenCombineitemQueryMap)
}

// ReleaseTaobaoQimenCombineitemQueryMap 释放TaobaoQimenCombineitemQueryMap
func ReleaseTaobaoQimenCombineitemQueryMap(v *TaobaoQimenCombineitemQueryMap) {
	poolTaobaoQimenCombineitemQueryMap.Put(v)
}
