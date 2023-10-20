package icbuproduct

import (
	"sync"
)

// ProductSupportTypeDto 结构体
type ProductSupportTypeDto struct {
	// 是否支持下单品
	SupportPostWholeSale bool `json:"support_post_whole_sale,omitempty" xml:"support_post_whole_sale,omitempty"`
	// 是否支持询盘品
	SupportPostSourcing bool `json:"support_post_sourcing,omitempty" xml:"support_post_sourcing,omitempty"`
}

var poolProductSupportTypeDto = sync.Pool{
	New: func() any {
		return new(ProductSupportTypeDto)
	},
}

// GetProductSupportTypeDto() 从对象池中获取ProductSupportTypeDto
func GetProductSupportTypeDto() *ProductSupportTypeDto {
	return poolProductSupportTypeDto.Get().(*ProductSupportTypeDto)
}

// ReleaseProductSupportTypeDto 释放ProductSupportTypeDto
func ReleaseProductSupportTypeDto(v *ProductSupportTypeDto) {
	v.SupportPostWholeSale = false
	v.SupportPostSourcing = false
	poolProductSupportTypeDto.Put(v)
}
