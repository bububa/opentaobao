package qimen

import (
	"sync"
)

// TaobaoQimenExpressinfoQueryMap 结构体
type TaobaoQimenExpressinfoQueryMap struct {
}

var poolTaobaoQimenExpressinfoQueryMap = sync.Pool{
	New: func() any {
		return new(TaobaoQimenExpressinfoQueryMap)
	},
}

// GetTaobaoQimenExpressinfoQueryMap() 从对象池中获取TaobaoQimenExpressinfoQueryMap
func GetTaobaoQimenExpressinfoQueryMap() *TaobaoQimenExpressinfoQueryMap {
	return poolTaobaoQimenExpressinfoQueryMap.Get().(*TaobaoQimenExpressinfoQueryMap)
}

// ReleaseTaobaoQimenExpressinfoQueryMap 释放TaobaoQimenExpressinfoQueryMap
func ReleaseTaobaoQimenExpressinfoQueryMap(v *TaobaoQimenExpressinfoQueryMap) {
	poolTaobaoQimenExpressinfoQueryMap.Put(v)
}
