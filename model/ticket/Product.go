package ticket

import (
	"sync"
)

// Product 结构体
type Product struct {
	// 标准收费项目ID
	AliProductId string `json:"ali_product_id,omitempty" xml:"ali_product_id,omitempty"`
	// 标准收费项目名称
	AliProductName string `json:"ali_product_name,omitempty" xml:"ali_product_name,omitempty"`
	// 商家收费项目ID
	OutProductId string `json:"out_product_id,omitempty" xml:"out_product_id,omitempty"`
	// 商家收费项目名称
	OutProductName string `json:"out_product_name,omitempty" xml:"out_product_name,omitempty"`
	// 商品ID
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 商品名称
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
}

var poolProduct = sync.Pool{
	New: func() any {
		return new(Product)
	},
}

// GetProduct() 从对象池中获取Product
func GetProduct() *Product {
	return poolProduct.Get().(*Product)
}

// ReleaseProduct 释放Product
func ReleaseProduct(v *Product) {
	v.AliProductId = ""
	v.AliProductName = ""
	v.OutProductId = ""
	v.OutProductName = ""
	v.ItemId = ""
	v.ItemName = ""
	poolProduct.Put(v)
}
