package xhotelonlineorder

import (
	"sync"
)

// OrderPromotionDo 结构体
type OrderPromotionDo struct {
	// OrderPromotionDetail
	CurrentOrderPromotion *OrderPromotionDetail `json:"current_order_promotion,omitempty" xml:"current_order_promotion,omitempty"`
}

var poolOrderPromotionDo = sync.Pool{
	New: func() any {
		return new(OrderPromotionDo)
	},
}

// GetOrderPromotionDo() 从对象池中获取OrderPromotionDo
func GetOrderPromotionDo() *OrderPromotionDo {
	return poolOrderPromotionDo.Get().(*OrderPromotionDo)
}

// ReleaseOrderPromotionDo 释放OrderPromotionDo
func ReleaseOrderPromotionDo(v *OrderPromotionDo) {
	v.CurrentOrderPromotion = nil
	poolOrderPromotionDo.Put(v)
}
