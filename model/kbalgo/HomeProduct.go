package kbalgo

import (
	"sync"
)

// HomeProduct 结构体
type HomeProduct struct {
	// 商品信息
	ProductInfos []ProductInfo `json:"product_infos,omitempty" xml:"product_infos>product_info,omitempty"`
	// 到家基本信息
	BaseInfo *BaseInfo `json:"base_info,omitempty" xml:"base_info,omitempty"`
}

var poolHomeProduct = sync.Pool{
	New: func() any {
		return new(HomeProduct)
	},
}

// GetHomeProduct() 从对象池中获取HomeProduct
func GetHomeProduct() *HomeProduct {
	return poolHomeProduct.Get().(*HomeProduct)
}

// ReleaseHomeProduct 释放HomeProduct
func ReleaseHomeProduct(v *HomeProduct) {
	v.ProductInfos = v.ProductInfos[:0]
	v.BaseInfo = nil
	poolHomeProduct.Put(v)
}
