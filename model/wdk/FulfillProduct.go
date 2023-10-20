package wdk

import (
	"sync"
)

// FulfillProduct 结构体
type FulfillProduct struct {
	// 商品编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 商品履约线路
	LineInstances string `json:"line_instances,omitempty" xml:"line_instances,omitempty"`
}

var poolFulfillProduct = sync.Pool{
	New: func() any {
		return new(FulfillProduct)
	},
}

// GetFulfillProduct() 从对象池中获取FulfillProduct
func GetFulfillProduct() *FulfillProduct {
	return poolFulfillProduct.Get().(*FulfillProduct)
}

// ReleaseFulfillProduct 释放FulfillProduct
func ReleaseFulfillProduct(v *FulfillProduct) {
	v.SkuCode = ""
	v.LineInstances = ""
	poolFulfillProduct.Put(v)
}
