package tmallsc

import (
	"sync"
)

// ServiceProduct 结构体
type ServiceProduct struct {
	// 服务产品名称
	ProductTitle string `json:"product_title,omitempty" xml:"product_title,omitempty"`
	// 服务产品id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
}

var poolServiceProduct = sync.Pool{
	New: func() any {
		return new(ServiceProduct)
	},
}

// GetServiceProduct() 从对象池中获取ServiceProduct
func GetServiceProduct() *ServiceProduct {
	return poolServiceProduct.Get().(*ServiceProduct)
}

// ReleaseServiceProduct 释放ServiceProduct
func ReleaseServiceProduct(v *ServiceProduct) {
	v.ProductTitle = ""
	v.ProductId = 0
	poolServiceProduct.Put(v)
}
