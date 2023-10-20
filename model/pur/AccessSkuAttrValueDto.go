package pur

import (
	"sync"
)

// AccessSkuAttrValueDto 结构体
type AccessSkuAttrValueDto struct {
	// 属性名称
	AttrName string `json:"attr_name,omitempty" xml:"attr_name,omitempty"`
	// 属性值名称
	AttrValueName string `json:"attr_value_name,omitempty" xml:"attr_value_name,omitempty"`
}

var poolAccessSkuAttrValueDto = sync.Pool{
	New: func() any {
		return new(AccessSkuAttrValueDto)
	},
}

// GetAccessSkuAttrValueDto() 从对象池中获取AccessSkuAttrValueDto
func GetAccessSkuAttrValueDto() *AccessSkuAttrValueDto {
	return poolAccessSkuAttrValueDto.Get().(*AccessSkuAttrValueDto)
}

// ReleaseAccessSkuAttrValueDto 释放AccessSkuAttrValueDto
func ReleaseAccessSkuAttrValueDto(v *AccessSkuAttrValueDto) {
	v.AttrName = ""
	v.AttrValueName = ""
	poolAccessSkuAttrValueDto.Put(v)
}
