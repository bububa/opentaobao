package alsc

import (
	"sync"
)

// PropertyValueTopDto 结构体
type PropertyValueTopDto struct {
	// 属性值
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// 属性ID
	PropertyId int64 `json:"property_id,omitempty" xml:"property_id,omitempty"`
}

var poolPropertyValueTopDto = sync.Pool{
	New: func() any {
		return new(PropertyValueTopDto)
	},
}

// GetPropertyValueTopDto() 从对象池中获取PropertyValueTopDto
func GetPropertyValueTopDto() *PropertyValueTopDto {
	return poolPropertyValueTopDto.Get().(*PropertyValueTopDto)
}

// ReleasePropertyValueTopDto 释放PropertyValueTopDto
func ReleasePropertyValueTopDto(v *PropertyValueTopDto) {
	v.Value = ""
	v.PropertyId = 0
	poolPropertyValueTopDto.Put(v)
}
