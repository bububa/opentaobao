package btrip

import (
	"sync"
)

// BaggageSubItem 结构体
type BaggageSubItem struct {
}

var poolBaggageSubItem = sync.Pool{
	New: func() any {
		return new(BaggageSubItem)
	},
}

// GetBaggageSubItem() 从对象池中获取BaggageSubItem
func GetBaggageSubItem() *BaggageSubItem {
	return poolBaggageSubItem.Get().(*BaggageSubItem)
}

// ReleaseBaggageSubItem 释放BaggageSubItem
func ReleaseBaggageSubItem(v *BaggageSubItem) {
	poolBaggageSubItem.Put(v)
}
