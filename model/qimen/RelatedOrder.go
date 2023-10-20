package qimen

import (
	"sync"
)

// RelatedOrder 结构体
type RelatedOrder struct {
	// 关联的订单类型(CG=采购单;DB=调拨单;CK=出库单;RK=入库单;只传英文编码)
	OrderType string `json:"orderType,omitempty" xml:"orderType,omitempty"`
	// 关联的订单编号
	OrderCode string `json:"orderCode,omitempty" xml:"orderCode,omitempty"`
}

var poolRelatedOrder = sync.Pool{
	New: func() any {
		return new(RelatedOrder)
	},
}

// GetRelatedOrder() 从对象池中获取RelatedOrder
func GetRelatedOrder() *RelatedOrder {
	return poolRelatedOrder.Get().(*RelatedOrder)
}

// ReleaseRelatedOrder 释放RelatedOrder
func ReleaseRelatedOrder(v *RelatedOrder) {
	v.OrderType = ""
	v.OrderCode = ""
	poolRelatedOrder.Put(v)
}
