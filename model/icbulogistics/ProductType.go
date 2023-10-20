package icbulogistics

import (
	"sync"
)

// ProductType 结构体
type ProductType struct {
	// 商品特性列表对象
	Children []Children `json:"children,omitempty" xml:"children>children,omitempty"`
	// 商品类型code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 商品类型
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

var poolProductType = sync.Pool{
	New: func() any {
		return new(ProductType)
	},
}

// GetProductType() 从对象池中获取ProductType
func GetProductType() *ProductType {
	return poolProductType.Get().(*ProductType)
}

// ReleaseProductType 释放ProductType
func ReleaseProductType(v *ProductType) {
	v.Children = v.Children[:0]
	v.Code = ""
	v.Name = ""
	poolProductType.Put(v)
}
