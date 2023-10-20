package ascpchannel

import (
	"sync"
)

// Products 结构体
type Products struct {
	// 经营模式
	SalesModes []string `json:"sales_modes,omitempty" xml:"sales_modes>string,omitempty"`
	// 图片
	Pictures []string `json:"pictures,omitempty" xml:"pictures>string,omitempty"`
	// 产品标题
	ProductTitle string `json:"product_title,omitempty" xml:"product_title,omitempty"`
	// 产品id
	ProductId string `json:"product_id,omitempty" xml:"product_id,omitempty"`
}

var poolProducts = sync.Pool{
	New: func() any {
		return new(Products)
	},
}

// GetProducts() 从对象池中获取Products
func GetProducts() *Products {
	return poolProducts.Get().(*Products)
}

// ReleaseProducts 释放Products
func ReleaseProducts(v *Products) {
	v.SalesModes = v.SalesModes[:0]
	v.Pictures = v.Pictures[:0]
	v.ProductTitle = ""
	v.ProductId = ""
	poolProducts.Put(v)
}
