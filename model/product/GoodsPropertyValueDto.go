package product

import (
	"sync"
)

// GoodsPropertyValueDto 结构体
type GoodsPropertyValueDto struct {
	// 属性值
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// 属性值id，可不传
	ValueId int64 `json:"value_id,omitempty" xml:"value_id,omitempty"`
	// 属性id
	PropertyId int64 `json:"property_id,omitempty" xml:"property_id,omitempty"`
}

var poolGoodsPropertyValueDto = sync.Pool{
	New: func() any {
		return new(GoodsPropertyValueDto)
	},
}

// GetGoodsPropertyValueDto() 从对象池中获取GoodsPropertyValueDto
func GetGoodsPropertyValueDto() *GoodsPropertyValueDto {
	return poolGoodsPropertyValueDto.Get().(*GoodsPropertyValueDto)
}

// ReleaseGoodsPropertyValueDto 释放GoodsPropertyValueDto
func ReleaseGoodsPropertyValueDto(v *GoodsPropertyValueDto) {
	v.Value = ""
	v.ValueId = 0
	v.PropertyId = 0
	poolGoodsPropertyValueDto.Put(v)
}
