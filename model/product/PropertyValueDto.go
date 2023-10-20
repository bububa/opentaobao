package product

import (
	"sync"
)

// PropertyValueDto 结构体
type PropertyValueDto struct {
	// 子属性列表
	ChildPropertyKeyValueList []ChildPropertyKeyValueDto `json:"child_property_key_value_list,omitempty" xml:"child_property_key_value_list>child_property_key_value_dto,omitempty"`
	// 属性值
	Value string `json:"value,omitempty" xml:"value,omitempty"`
}

var poolPropertyValueDto = sync.Pool{
	New: func() any {
		return new(PropertyValueDto)
	},
}

// GetPropertyValueDto() 从对象池中获取PropertyValueDto
func GetPropertyValueDto() *PropertyValueDto {
	return poolPropertyValueDto.Get().(*PropertyValueDto)
}

// ReleasePropertyValueDto 释放PropertyValueDto
func ReleasePropertyValueDto(v *PropertyValueDto) {
	v.ChildPropertyKeyValueList = v.ChildPropertyKeyValueList[:0]
	v.Value = ""
	poolPropertyValueDto.Put(v)
}
