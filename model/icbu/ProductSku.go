package icbu

import (
	"sync"
)

// ProductSku 结构体
type ProductSku struct {
	// 商品属性
	Attributes []ProductAttribute `json:"attributes,omitempty" xml:"attributes>product_attribute,omitempty"`
	// 需要失效的SKU的详细定义
	ExcludeSkus []SkuDetail `json:"exclude_skus,omitempty" xml:"exclude_skus>sku_detail,omitempty"`
	// 单个SKU详细定义
	SpecialSkus []SkuDetail `json:"special_skus,omitempty" xml:"special_skus>sku_detail,omitempty"`
}

var poolProductSku = sync.Pool{
	New: func() any {
		return new(ProductSku)
	},
}

// GetProductSku() 从对象池中获取ProductSku
func GetProductSku() *ProductSku {
	return poolProductSku.Get().(*ProductSku)
}

// ReleaseProductSku 释放ProductSku
func ReleaseProductSku(v *ProductSku) {
	v.Attributes = v.Attributes[:0]
	v.ExcludeSkus = v.ExcludeSkus[:0]
	v.SpecialSkus = v.SpecialSkus[:0]
	poolProductSku.Put(v)
}
