package axintrade

import (
	"sync"
)

// ProductPriceDto 结构体
type ProductPriceDto struct {
	// 日期 格式yyyy-MM-dd
	Date string `json:"date,omitempty" xml:"date,omitempty"`
	// 价格，单位为分
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
}

var poolProductPriceDto = sync.Pool{
	New: func() any {
		return new(ProductPriceDto)
	},
}

// GetProductPriceDto() 从对象池中获取ProductPriceDto
func GetProductPriceDto() *ProductPriceDto {
	return poolProductPriceDto.Get().(*ProductPriceDto)
}

// ReleaseProductPriceDto 释放ProductPriceDto
func ReleaseProductPriceDto(v *ProductPriceDto) {
	v.Date = ""
	v.Price = 0
	poolProductPriceDto.Put(v)
}
