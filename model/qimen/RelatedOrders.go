package qimen

import (
	"sync"
)

// RelatedOrders 结构体
type RelatedOrders struct {
	// 关联的订单类型,关联的订单类型,string(50),,
	OrderType string `json:"orderType,omitempty" xml:"orderType,omitempty"`
	// 关联的订单编号,关联的订单编号,string(50),,
	OrderCode string `json:"orderCode,omitempty" xml:"orderCode,omitempty"`
}

var poolRelatedOrders = sync.Pool{
	New: func() any {
		return new(RelatedOrders)
	},
}

// GetRelatedOrders() 从对象池中获取RelatedOrders
func GetRelatedOrders() *RelatedOrders {
	return poolRelatedOrders.Get().(*RelatedOrders)
}

// ReleaseRelatedOrders 释放RelatedOrders
func ReleaseRelatedOrders(v *RelatedOrders) {
	v.OrderType = ""
	v.OrderCode = ""
	poolRelatedOrders.Put(v)
}
