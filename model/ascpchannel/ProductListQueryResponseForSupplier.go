package ascpchannel

import (
	"sync"
)

// ProductListQueryResponseForSupplier 结构体
type ProductListQueryResponseForSupplier struct {
	// 产品列表
	ProductList []ProductList `json:"product_list,omitempty" xml:"product_list>product_list,omitempty"`
	// 总数量
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

var poolProductListQueryResponseForSupplier = sync.Pool{
	New: func() any {
		return new(ProductListQueryResponseForSupplier)
	},
}

// GetProductListQueryResponseForSupplier() 从对象池中获取ProductListQueryResponseForSupplier
func GetProductListQueryResponseForSupplier() *ProductListQueryResponseForSupplier {
	return poolProductListQueryResponseForSupplier.Get().(*ProductListQueryResponseForSupplier)
}

// ReleaseProductListQueryResponseForSupplier 释放ProductListQueryResponseForSupplier
func ReleaseProductListQueryResponseForSupplier(v *ProductListQueryResponseForSupplier) {
	v.ProductList = v.ProductList[:0]
	v.TotalCount = 0
	poolProductListQueryResponseForSupplier.Put(v)
}
