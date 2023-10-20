package aedropshiper

import (
	"sync"
)

// ProductIdConverterResultDto 结构体
type ProductIdConverterResultDto struct {
	// sub productId
	SubProductId string `json:"sub_product_id,omitempty" xml:"sub_product_id,omitempty"`
	// main productId
	MainProductId int64 `json:"main_product_id,omitempty" xml:"main_product_id,omitempty"`
}

var poolProductIdConverterResultDto = sync.Pool{
	New: func() any {
		return new(ProductIdConverterResultDto)
	},
}

// GetProductIdConverterResultDto() 从对象池中获取ProductIdConverterResultDto
func GetProductIdConverterResultDto() *ProductIdConverterResultDto {
	return poolProductIdConverterResultDto.Get().(*ProductIdConverterResultDto)
}

// ReleaseProductIdConverterResultDto 释放ProductIdConverterResultDto
func ReleaseProductIdConverterResultDto(v *ProductIdConverterResultDto) {
	v.SubProductId = ""
	v.MainProductId = 0
	poolProductIdConverterResultDto.Put(v)
}
