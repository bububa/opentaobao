package aesolution

import (
	"sync"
)

// SkuValueSimplifiedInfoDto 结构体
type SkuValueSimplifiedInfoDto struct {
	// aliexpress sku value name
	AliexpressSkuValueName string `json:"aliexpress_sku_value_name,omitempty" xml:"aliexpress_sku_value_name,omitempty"`
}

var poolSkuValueSimplifiedInfoDto = sync.Pool{
	New: func() any {
		return new(SkuValueSimplifiedInfoDto)
	},
}

// GetSkuValueSimplifiedInfoDto() 从对象池中获取SkuValueSimplifiedInfoDto
func GetSkuValueSimplifiedInfoDto() *SkuValueSimplifiedInfoDto {
	return poolSkuValueSimplifiedInfoDto.Get().(*SkuValueSimplifiedInfoDto)
}

// ReleaseSkuValueSimplifiedInfoDto 释放SkuValueSimplifiedInfoDto
func ReleaseSkuValueSimplifiedInfoDto(v *SkuValueSimplifiedInfoDto) {
	v.AliexpressSkuValueName = ""
	poolSkuValueSimplifiedInfoDto.Put(v)
}
