package icbudropshipping

import (
	"sync"
)

// LadderPrice 结构体
type LadderPrice struct {
	// price
	Price float64 `json:"price,omitempty" xml:"price,omitempty"`
	// If it is -1, it means the maximum
	MaxQuantity int64 `json:"max_quantity,omitempty" xml:"max_quantity,omitempty"`
	// min quantity
	MinQuantity int64 `json:"min_quantity,omitempty" xml:"min_quantity,omitempty"`
}

var poolLadderPrice = sync.Pool{
	New: func() any {
		return new(LadderPrice)
	},
}

// GetLadderPrice() 从对象池中获取LadderPrice
func GetLadderPrice() *LadderPrice {
	return poolLadderPrice.Get().(*LadderPrice)
}

// ReleaseLadderPrice 释放LadderPrice
func ReleaseLadderPrice(v *LadderPrice) {
	v.Price = 0
	v.MaxQuantity = 0
	v.MinQuantity = 0
	poolLadderPrice.Put(v)
}
