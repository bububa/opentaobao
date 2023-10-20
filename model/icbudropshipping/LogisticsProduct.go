package icbudropshipping

import (
	"sync"
)

// LogisticsProduct 结构体
type LogisticsProduct struct {
	// Product Id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// quantity
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// Sku ID
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}

var poolLogisticsProduct = sync.Pool{
	New: func() any {
		return new(LogisticsProduct)
	},
}

// GetLogisticsProduct() 从对象池中获取LogisticsProduct
func GetLogisticsProduct() *LogisticsProduct {
	return poolLogisticsProduct.Get().(*LogisticsProduct)
}

// ReleaseLogisticsProduct 释放LogisticsProduct
func ReleaseLogisticsProduct(v *LogisticsProduct) {
	v.ProductId = 0
	v.Quantity = 0
	v.SkuId = 0
	poolLogisticsProduct.Put(v)
}
