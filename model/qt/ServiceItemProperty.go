package qt

import (
	"sync"
)

// ServiceItemProperty 结构体
type ServiceItemProperty struct {
	// 属性列表
	ItemPropertyValues []ItemPropertyValues `json:"item_property_values,omitempty" xml:"item_property_values>item_property_values,omitempty"`
	// 服务名称
	ServiceName string `json:"service_name,omitempty" xml:"service_name,omitempty"`
	// 服务收费项名称
	ServiceItemName string `json:"service_item_name,omitempty" xml:"service_item_name,omitempty"`
	// 服务收费项代码
	ServiceItemCode string `json:"service_item_code,omitempty" xml:"service_item_code,omitempty"`
	// 基础价格
	BasicPrice string `json:"basic_price,omitempty" xml:"basic_price,omitempty"`
	// 服务商的nick
	Nick string `json:"nick,omitempty" xml:"nick,omitempty"`
	// 质检服务简介
	Description string `json:"description,omitempty" xml:"description,omitempty"`
}

var poolServiceItemProperty = sync.Pool{
	New: func() any {
		return new(ServiceItemProperty)
	},
}

// GetServiceItemProperty() 从对象池中获取ServiceItemProperty
func GetServiceItemProperty() *ServiceItemProperty {
	return poolServiceItemProperty.Get().(*ServiceItemProperty)
}

// ReleaseServiceItemProperty 释放ServiceItemProperty
func ReleaseServiceItemProperty(v *ServiceItemProperty) {
	v.ItemPropertyValues = v.ItemPropertyValues[:0]
	v.ServiceName = ""
	v.ServiceItemName = ""
	v.ServiceItemCode = ""
	v.BasicPrice = ""
	v.Nick = ""
	v.Description = ""
	poolServiceItemProperty.Put(v)
}
