package aesolution

import (
	"sync"
)

// SynchronizeSkuRequestDto 结构体
type SynchronizeSkuRequestDto struct {
	// sku code
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// price of an sku
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// discount_price of an sku. If not set, the discount_price will be erased.
	DiscountPrice string `json:"discount_price,omitempty" xml:"discount_price,omitempty"`
	// inventory
	Inventory int64 `json:"inventory,omitempty" xml:"inventory,omitempty"`
}

var poolSynchronizeSkuRequestDto = sync.Pool{
	New: func() any {
		return new(SynchronizeSkuRequestDto)
	},
}

// GetSynchronizeSkuRequestDto() 从对象池中获取SynchronizeSkuRequestDto
func GetSynchronizeSkuRequestDto() *SynchronizeSkuRequestDto {
	return poolSynchronizeSkuRequestDto.Get().(*SynchronizeSkuRequestDto)
}

// ReleaseSynchronizeSkuRequestDto 释放SynchronizeSkuRequestDto
func ReleaseSynchronizeSkuRequestDto(v *SynchronizeSkuRequestDto) {
	v.SkuCode = ""
	v.Price = ""
	v.DiscountPrice = ""
	v.Inventory = 0
	poolSynchronizeSkuRequestDto.Put(v)
}
