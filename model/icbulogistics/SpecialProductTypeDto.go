package icbulogistics

import (
	"sync"
)

// SpecialProductTypeDto 结构体
type SpecialProductTypeDto struct {
	// 列表对象
	Childrens []SpecialProductTypeDto `json:"childrens,omitempty" xml:"childrens>special_product_type_dto,omitempty"`
	// 商品类型code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 商品类型
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

var poolSpecialProductTypeDto = sync.Pool{
	New: func() any {
		return new(SpecialProductTypeDto)
	},
}

// GetSpecialProductTypeDto() 从对象池中获取SpecialProductTypeDto
func GetSpecialProductTypeDto() *SpecialProductTypeDto {
	return poolSpecialProductTypeDto.Get().(*SpecialProductTypeDto)
}

// ReleaseSpecialProductTypeDto 释放SpecialProductTypeDto
func ReleaseSpecialProductTypeDto(v *SpecialProductTypeDto) {
	v.Childrens = v.Childrens[:0]
	v.Code = ""
	v.Name = ""
	poolSpecialProductTypeDto.Put(v)
}
