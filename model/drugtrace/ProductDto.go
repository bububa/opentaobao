package drugtrace

import (
	"sync"
)

// ProductDto 结构体
type ProductDto struct {
	// 子类
	SubTypeList []SubType `json:"sub_type_list,omitempty" xml:"sub_type_list>sub_type,omitempty"`
	// 产品 码
	ProductCode string `json:"product_code,omitempty" xml:"product_code,omitempty"`
	// 备注
	Comment string `json:"comment,omitempty" xml:"comment,omitempty"`
	// 药品名称
	ProductName string `json:"product_name,omitempty" xml:"product_name,omitempty"`
}

var poolProductDto = sync.Pool{
	New: func() any {
		return new(ProductDto)
	},
}

// GetProductDto() 从对象池中获取ProductDto
func GetProductDto() *ProductDto {
	return poolProductDto.Get().(*ProductDto)
}

// ReleaseProductDto 释放ProductDto
func ReleaseProductDto(v *ProductDto) {
	v.SubTypeList = v.SubTypeList[:0]
	v.ProductCode = ""
	v.Comment = ""
	v.ProductName = ""
	poolProductDto.Put(v)
}
