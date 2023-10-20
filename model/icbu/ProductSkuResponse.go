package icbu

import (
	"sync"
)

// ProductSkuResponse 结构体
type ProductSkuResponse struct {
	// SKU使用的属性
	SkuAttributes []SkuAttribute `json:"sku_attributes,omitempty" xml:"sku_attributes>sku_attribute,omitempty"`
	// SKU定义
	Skus []SkuDefinition `json:"skus,omitempty" xml:"skus>sku_definition,omitempty"`
}

var poolProductSkuResponse = sync.Pool{
	New: func() any {
		return new(ProductSkuResponse)
	},
}

// GetProductSkuResponse() 从对象池中获取ProductSkuResponse
func GetProductSkuResponse() *ProductSkuResponse {
	return poolProductSkuResponse.Get().(*ProductSkuResponse)
}

// ReleaseProductSkuResponse 释放ProductSkuResponse
func ReleaseProductSkuResponse(v *ProductSkuResponse) {
	v.SkuAttributes = v.SkuAttributes[:0]
	v.Skus = v.Skus[:0]
	poolProductSkuResponse.Put(v)
}
