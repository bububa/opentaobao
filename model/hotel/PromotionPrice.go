package hotel

import (
	"sync"
)

// PromotionPrice 结构体
type PromotionPrice struct {
}

var poolPromotionPrice = sync.Pool{
	New: func() any {
		return new(PromotionPrice)
	},
}

// GetPromotionPrice() 从对象池中获取PromotionPrice
func GetPromotionPrice() *PromotionPrice {
	return poolPromotionPrice.Get().(*PromotionPrice)
}

// ReleasePromotionPrice 释放PromotionPrice
func ReleasePromotionPrice(v *PromotionPrice) {
	poolPromotionPrice.Put(v)
}
