package icbu

import (
	"sync"
)

// BulkDiscountPrice 结构体
type BulkDiscountPrice struct {
	// 价格，范围是0.01-9999999.00
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 起始数量，范围是1-99999
	StartQuantity int64 `json:"start_quantity,omitempty" xml:"start_quantity,omitempty"`
}

var poolBulkDiscountPrice = sync.Pool{
	New: func() any {
		return new(BulkDiscountPrice)
	},
}

// GetBulkDiscountPrice() 从对象池中获取BulkDiscountPrice
func GetBulkDiscountPrice() *BulkDiscountPrice {
	return poolBulkDiscountPrice.Get().(*BulkDiscountPrice)
}

// ReleaseBulkDiscountPrice 释放BulkDiscountPrice
func ReleaseBulkDiscountPrice(v *BulkDiscountPrice) {
	v.Price = ""
	v.StartQuantity = 0
	poolBulkDiscountPrice.Put(v)
}
