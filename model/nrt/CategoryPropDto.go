package nrt

import (
	"sync"
)

// CategoryPropDto 结构体
type CategoryPropDto struct {
	// 属性名
	PropertyName string `json:"property_name,omitempty" xml:"property_name,omitempty"`
	// 值名
	ValueText string `json:"value_text,omitempty" xml:"value_text,omitempty"`
	// 属性ID
	PropertyId int64 `json:"property_id,omitempty" xml:"property_id,omitempty"`
	// 值ID
	ValueId int64 `json:"value_id,omitempty" xml:"value_id,omitempty"`
}

var poolCategoryPropDto = sync.Pool{
	New: func() any {
		return new(CategoryPropDto)
	},
}

// GetCategoryPropDto() 从对象池中获取CategoryPropDto
func GetCategoryPropDto() *CategoryPropDto {
	return poolCategoryPropDto.Get().(*CategoryPropDto)
}

// ReleaseCategoryPropDto 释放CategoryPropDto
func ReleaseCategoryPropDto(v *CategoryPropDto) {
	v.PropertyName = ""
	v.ValueText = ""
	v.PropertyId = 0
	v.ValueId = 0
	poolCategoryPropDto.Put(v)
}
