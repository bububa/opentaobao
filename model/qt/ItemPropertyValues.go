package qt

import (
	"sync"
)

// ItemPropertyValues 结构体
type ItemPropertyValues struct {
	// 属性值列表.
	PropertyValues []string `json:"property_values,omitempty" xml:"property_values>string,omitempty"`
	// 属性名称
	PropertyName string `json:"property_name,omitempty" xml:"property_name,omitempty"`
	// 服务属性id
	PropertyId int64 `json:"property_id,omitempty" xml:"property_id,omitempty"`
}

var poolItemPropertyValues = sync.Pool{
	New: func() any {
		return new(ItemPropertyValues)
	},
}

// GetItemPropertyValues() 从对象池中获取ItemPropertyValues
func GetItemPropertyValues() *ItemPropertyValues {
	return poolItemPropertyValues.Get().(*ItemPropertyValues)
}

// ReleaseItemPropertyValues 释放ItemPropertyValues
func ReleaseItemPropertyValues(v *ItemPropertyValues) {
	v.PropertyValues = v.PropertyValues[:0]
	v.PropertyName = ""
	v.PropertyId = 0
	poolItemPropertyValues.Put(v)
}
