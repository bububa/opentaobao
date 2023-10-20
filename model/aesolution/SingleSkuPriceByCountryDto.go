package aesolution

import (
	"sync"
)

// SingleSkuPriceByCountryDto 结构体
type SingleSkuPriceByCountryDto struct {
	// sku_code, must existed in  multiple_sku_update_list
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// Value of price configuration. If the price of a specific country is set, it must be greater than or equal to 70% of the original sku price in multiple_sku_update_list
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// Value of discount_price for each country
	DiscountPrice string `json:"discount_price,omitempty" xml:"discount_price,omitempty"`
	// Deprecated. Please do not use.
	Value string `json:"value,omitempty" xml:"value,omitempty"`
}

var poolSingleSkuPriceByCountryDto = sync.Pool{
	New: func() any {
		return new(SingleSkuPriceByCountryDto)
	},
}

// GetSingleSkuPriceByCountryDto() 从对象池中获取SingleSkuPriceByCountryDto
func GetSingleSkuPriceByCountryDto() *SingleSkuPriceByCountryDto {
	return poolSingleSkuPriceByCountryDto.Get().(*SingleSkuPriceByCountryDto)
}

// ReleaseSingleSkuPriceByCountryDto 释放SingleSkuPriceByCountryDto
func ReleaseSingleSkuPriceByCountryDto(v *SingleSkuPriceByCountryDto) {
	v.SkuCode = ""
	v.Price = ""
	v.DiscountPrice = ""
	v.Value = ""
	poolSingleSkuPriceByCountryDto.Put(v)
}
