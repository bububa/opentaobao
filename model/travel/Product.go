package travel

import (
	"sync"
)

// Product 结构体
type Product struct {
	// 描述
	Descr string `json:"descr,omitempty" xml:"descr,omitempty"`
	// 资源元素的外部商家编码
	ElementId string `json:"element_id,omitempty" xml:"element_id,omitempty"`
	// 关联的套餐id
	PackageId int64 `json:"package_id,omitempty" xml:"package_id,omitempty"`
	// 数量
	Num int64 `json:"num,omitempty" xml:"num,omitempty"`
	// 是否主元素
	MainProduct bool `json:"main_product,omitempty" xml:"main_product,omitempty"`
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
	v.Descr = ""
	v.ElementId = ""
	v.PackageId = 0
	v.Num = 0
	v.MainProduct = false
	poolProduct.Put(v)
}
