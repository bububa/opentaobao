package product

import (
	"sync"
)

// MultiPropertyValueDto 结构体
type MultiPropertyValueDto struct {
	// 属性值对象数组
	PropertyValueList []PropertyValueDto `json:"property_value_list,omitempty" xml:"property_value_list>property_value_dto,omitempty"`
	// 父属性名
	PropertyName string `json:"property_name,omitempty" xml:"property_name,omitempty"`
}

var poolMultiPropertyValueDto = sync.Pool{
	New: func() any {
		return new(MultiPropertyValueDto)
	},
}

// GetMultiPropertyValueDto() 从对象池中获取MultiPropertyValueDto
func GetMultiPropertyValueDto() *MultiPropertyValueDto {
	return poolMultiPropertyValueDto.Get().(*MultiPropertyValueDto)
}

// ReleaseMultiPropertyValueDto 释放MultiPropertyValueDto
func ReleaseMultiPropertyValueDto(v *MultiPropertyValueDto) {
	v.PropertyValueList = v.PropertyValueList[:0]
	v.PropertyName = ""
	poolMultiPropertyValueDto.Put(v)
}
