package util

import (
	"sync"
)

// PriceTopDto 结构体
type PriceTopDto struct {
	// 溢价信息，范围5-300，不在范围内会默认设置为5
	Discount int64 `json:"discount,omitempty" xml:"discount,omitempty"`
}

var poolPriceTopDto = sync.Pool{
	New: func() any {
		return new(PriceTopDto)
	},
}

// GetPriceTopDto() 从对象池中获取PriceTopDto
func GetPriceTopDto() *PriceTopDto {
	return poolPriceTopDto.Get().(*PriceTopDto)
}

// ReleasePriceTopDto 释放PriceTopDto
func ReleasePriceTopDto(v *PriceTopDto) {
	v.Discount = 0
	poolPriceTopDto.Put(v)
}
