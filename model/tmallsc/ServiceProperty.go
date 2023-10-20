package tmallsc

import (
	"sync"
)

// ServiceProperty 结构体
type ServiceProperty struct {
	// 服务属性值枚举
	ServicePropertyValueEnum []string `json:"service_property_value_enum,omitempty" xml:"service_property_value_enum>string,omitempty"`
	// 服务属性名称
	ServicePropertyName string `json:"service_property_name,omitempty" xml:"service_property_name,omitempty"`
}

var poolServiceProperty = sync.Pool{
	New: func() any {
		return new(ServiceProperty)
	},
}

// GetServiceProperty() 从对象池中获取ServiceProperty
func GetServiceProperty() *ServiceProperty {
	return poolServiceProperty.Get().(*ServiceProperty)
}

// ReleaseServiceProperty 释放ServiceProperty
func ReleaseServiceProperty(v *ServiceProperty) {
	v.ServicePropertyValueEnum = v.ServicePropertyValueEnum[:0]
	v.ServicePropertyName = ""
	poolServiceProperty.Put(v)
}
