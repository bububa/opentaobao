package wms

import (
	"sync"
)

// Orderitemlistwlbwmsstockoutordernotify 结构体
type Orderitemlistwlbwmsstockoutordernotify struct {
	// 订单商品信息
	OrderItem *Orderitemwlbwmsstockoutordernotify `json:"order_item,omitempty" xml:"order_item,omitempty"`
}

var poolOrderitemlistwlbwmsstockoutordernotify = sync.Pool{
	New: func() any {
		return new(Orderitemlistwlbwmsstockoutordernotify)
	},
}

// GetOrderitemlistwlbwmsstockoutordernotify() 从对象池中获取Orderitemlistwlbwmsstockoutordernotify
func GetOrderitemlistwlbwmsstockoutordernotify() *Orderitemlistwlbwmsstockoutordernotify {
	return poolOrderitemlistwlbwmsstockoutordernotify.Get().(*Orderitemlistwlbwmsstockoutordernotify)
}

// ReleaseOrderitemlistwlbwmsstockoutordernotify 释放Orderitemlistwlbwmsstockoutordernotify
func ReleaseOrderitemlistwlbwmsstockoutordernotify(v *Orderitemlistwlbwmsstockoutordernotify) {
	v.OrderItem = nil
	poolOrderitemlistwlbwmsstockoutordernotify.Put(v)
}
