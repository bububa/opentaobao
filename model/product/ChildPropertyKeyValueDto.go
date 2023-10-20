package product

import (
	"sync"
)

// ChildPropertyKeyValueDto 结构体
type ChildPropertyKeyValueDto struct {
	// 属性键
	Key string `json:"key,omitempty" xml:"key,omitempty"`
	// 属性值
	Value string `json:"value,omitempty" xml:"value,omitempty"`
}

var poolChildPropertyKeyValueDto = sync.Pool{
	New: func() any {
		return new(ChildPropertyKeyValueDto)
	},
}

// GetChildPropertyKeyValueDto() 从对象池中获取ChildPropertyKeyValueDto
func GetChildPropertyKeyValueDto() *ChildPropertyKeyValueDto {
	return poolChildPropertyKeyValueDto.Get().(*ChildPropertyKeyValueDto)
}

// ReleaseChildPropertyKeyValueDto 释放ChildPropertyKeyValueDto
func ReleaseChildPropertyKeyValueDto(v *ChildPropertyKeyValueDto) {
	v.Key = ""
	v.Value = ""
	poolChildPropertyKeyValueDto.Put(v)
}
