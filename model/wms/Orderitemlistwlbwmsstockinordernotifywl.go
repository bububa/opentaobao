package wms

import (
	"sync"
)

// Orderitemlistwlbwmsstockinordernotifywl 结构体
type Orderitemlistwlbwmsstockinordernotifywl struct {
	// 系统自动生成
	OrderItem *Orderitemwlbwmsstockinordernotifywl `json:"order_item,omitempty" xml:"order_item,omitempty"`
}

var poolOrderitemlistwlbwmsstockinordernotifywl = sync.Pool{
	New: func() any {
		return new(Orderitemlistwlbwmsstockinordernotifywl)
	},
}

// GetOrderitemlistwlbwmsstockinordernotifywl() 从对象池中获取Orderitemlistwlbwmsstockinordernotifywl
func GetOrderitemlistwlbwmsstockinordernotifywl() *Orderitemlistwlbwmsstockinordernotifywl {
	return poolOrderitemlistwlbwmsstockinordernotifywl.Get().(*Orderitemlistwlbwmsstockinordernotifywl)
}

// ReleaseOrderitemlistwlbwmsstockinordernotifywl 释放Orderitemlistwlbwmsstockinordernotifywl
func ReleaseOrderitemlistwlbwmsstockinordernotifywl(v *Orderitemlistwlbwmsstockinordernotifywl) {
	v.OrderItem = nil
	poolOrderitemlistwlbwmsstockinordernotifywl.Put(v)
}
