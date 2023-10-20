package wms

import (
	"sync"
)

// Orderitemlist 结构体
type Orderitemlist struct {
	// 订单商品信息
	OrderItem *Orderitem `json:"order_item,omitempty" xml:"order_item,omitempty"`
}

var poolOrderitemlist = sync.Pool{
	New: func() any {
		return new(Orderitemlist)
	},
}

// GetOrderitemlist() 从对象池中获取Orderitemlist
func GetOrderitemlist() *Orderitemlist {
	return poolOrderitemlist.Get().(*Orderitemlist)
}

// ReleaseOrderitemlist 释放Orderitemlist
func ReleaseOrderitemlist(v *Orderitemlist) {
	v.OrderItem = nil
	poolOrderitemlist.Put(v)
}
