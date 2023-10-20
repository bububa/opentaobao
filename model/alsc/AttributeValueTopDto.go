package alsc

import (
	"sync"
)

// AttributeValueTopDto 结构体
type AttributeValueTopDto struct {
	// 属性Value
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// 属性KEY
	Key string `json:"key,omitempty" xml:"key,omitempty"`
}

var poolAttributeValueTopDto = sync.Pool{
	New: func() any {
		return new(AttributeValueTopDto)
	},
}

// GetAttributeValueTopDto() 从对象池中获取AttributeValueTopDto
func GetAttributeValueTopDto() *AttributeValueTopDto {
	return poolAttributeValueTopDto.Get().(*AttributeValueTopDto)
}

// ReleaseAttributeValueTopDto 释放AttributeValueTopDto
func ReleaseAttributeValueTopDto(v *AttributeValueTopDto) {
	v.Value = ""
	v.Key = ""
	poolAttributeValueTopDto.Put(v)
}
