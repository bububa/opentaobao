package ascpchannel

import (
	"sync"
)

// ProductSkuLinkDto 结构体
type ProductSkuLinkDto struct {
	// 分销商商品 skuid
	OutSkuId string `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
	// 供应商产品 skuId
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}

var poolProductSkuLinkDto = sync.Pool{
	New: func() any {
		return new(ProductSkuLinkDto)
	},
}

// GetProductSkuLinkDto() 从对象池中获取ProductSkuLinkDto
func GetProductSkuLinkDto() *ProductSkuLinkDto {
	return poolProductSkuLinkDto.Get().(*ProductSkuLinkDto)
}

// ReleaseProductSkuLinkDto 释放ProductSkuLinkDto
func ReleaseProductSkuLinkDto(v *ProductSkuLinkDto) {
	v.OutSkuId = ""
	v.SkuId = ""
	poolProductSkuLinkDto.Put(v)
}
