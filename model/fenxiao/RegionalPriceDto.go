package fenxiao

import (
	"sync"
)

// RegionalPriceDto 结构体
type RegionalPriceDto struct {
	// 市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 省
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 区县，特殊可选
	District string `json:"district,omitempty" xml:"district,omitempty"`
	// 街道，特殊可选
	Street string `json:"street,omitempty" xml:"street,omitempty"`
	// 金额（分）
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
}

var poolRegionalPriceDto = sync.Pool{
	New: func() any {
		return new(RegionalPriceDto)
	},
}

// GetRegionalPriceDto() 从对象池中获取RegionalPriceDto
func GetRegionalPriceDto() *RegionalPriceDto {
	return poolRegionalPriceDto.Get().(*RegionalPriceDto)
}

// ReleaseRegionalPriceDto 释放RegionalPriceDto
func ReleaseRegionalPriceDto(v *RegionalPriceDto) {
	v.City = ""
	v.Province = ""
	v.District = ""
	v.Street = ""
	v.Price = 0
	poolRegionalPriceDto.Put(v)
}
