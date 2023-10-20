package wms

import (
	"sync"
)

// Orderitemlistwlbwmsreturnordernotify 结构体
type Orderitemlistwlbwmsreturnordernotify struct {
	// 1
	OrderItem *Orderitemwlbwmsreturnordernotify `json:"order_item,omitempty" xml:"order_item,omitempty"`
}

var poolOrderitemlistwlbwmsreturnordernotify = sync.Pool{
	New: func() any {
		return new(Orderitemlistwlbwmsreturnordernotify)
	},
}

// GetOrderitemlistwlbwmsreturnordernotify() 从对象池中获取Orderitemlistwlbwmsreturnordernotify
func GetOrderitemlistwlbwmsreturnordernotify() *Orderitemlistwlbwmsreturnordernotify {
	return poolOrderitemlistwlbwmsreturnordernotify.Get().(*Orderitemlistwlbwmsreturnordernotify)
}

// ReleaseOrderitemlistwlbwmsreturnordernotify 释放Orderitemlistwlbwmsreturnordernotify
func ReleaseOrderitemlistwlbwmsreturnordernotify(v *Orderitemlistwlbwmsreturnordernotify) {
	v.OrderItem = nil
	poolOrderitemlistwlbwmsreturnordernotify.Put(v)
}
