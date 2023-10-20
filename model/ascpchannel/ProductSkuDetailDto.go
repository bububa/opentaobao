package ascpchannel

import (
	"sync"
)

// ProductSkuDetailDto 结构体
type ProductSkuDetailDto struct {
	// sku 销售属性
	Properties []string `json:"properties,omitempty" xml:"properties>string,omitempty"`
	// sku图片链接
	Picture []string `json:"picture,omitempty" xml:"picture>string,omitempty"`
	// skuId
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}

var poolProductSkuDetailDto = sync.Pool{
	New: func() any {
		return new(ProductSkuDetailDto)
	},
}

// GetProductSkuDetailDto() 从对象池中获取ProductSkuDetailDto
func GetProductSkuDetailDto() *ProductSkuDetailDto {
	return poolProductSkuDetailDto.Get().(*ProductSkuDetailDto)
}

// ReleaseProductSkuDetailDto 释放ProductSkuDetailDto
func ReleaseProductSkuDetailDto(v *ProductSkuDetailDto) {
	v.Properties = v.Properties[:0]
	v.Picture = v.Picture[:0]
	v.SkuId = ""
	poolProductSkuDetailDto.Put(v)
}
