package btrip

import (
	"sync"
)

// RefundSubItem 结构体
type RefundSubItem struct {
}

var poolRefundSubItem = sync.Pool{
	New: func() any {
		return new(RefundSubItem)
	},
}

// GetRefundSubItem() 从对象池中获取RefundSubItem
func GetRefundSubItem() *RefundSubItem {
	return poolRefundSubItem.Get().(*RefundSubItem)
}

// ReleaseRefundSubItem 释放RefundSubItem
func ReleaseRefundSubItem(v *RefundSubItem) {
	poolRefundSubItem.Put(v)
}
