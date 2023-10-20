package tbtrade

import (
	"sync"
)

// OrderMerge 结构体
type OrderMerge struct {
	// 收件人ID (Open Addressee ID)，长度在128个字符之内。
	Oaid string `json:"oaid,omitempty" xml:"oaid,omitempty"`
	// 订单ID
	Tid string `json:"tid,omitempty" xml:"tid,omitempty"`
}

var poolOrderMerge = sync.Pool{
	New: func() any {
		return new(OrderMerge)
	},
}

// GetOrderMerge() 从对象池中获取OrderMerge
func GetOrderMerge() *OrderMerge {
	return poolOrderMerge.Get().(*OrderMerge)
}

// ReleaseOrderMerge 释放OrderMerge
func ReleaseOrderMerge(v *OrderMerge) {
	v.Oaid = ""
	v.Tid = ""
	poolOrderMerge.Put(v)
}
