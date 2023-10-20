package aesolution

import (
	"sync"
)

// GlobalAeopAeProductProperty 结构体
type GlobalAeopAeProductProperty struct {
	// Customized attribute name
	AttrName string `json:"attr_name,omitempty" xml:"attr_name,omitempty"`
	// Customized attribute value
	AttrValue string `json:"attr_value,omitempty" xml:"attr_value,omitempty"`
	// unit of customized attribute value
	AttrValueUnit string `json:"attr_value_unit,omitempty" xml:"attr_value_unit,omitempty"`
	// Attribute Name ID
	AttrNameId int64 `json:"attr_name_id,omitempty" xml:"attr_name_id,omitempty"`
	// Attribute Value ID
	AttrValueId int64 `json:"attr_value_id,omitempty" xml:"attr_value_id,omitempty"`
}

var poolGlobalAeopAeProductProperty = sync.Pool{
	New: func() any {
		return new(GlobalAeopAeProductProperty)
	},
}

// GetGlobalAeopAeProductProperty() 从对象池中获取GlobalAeopAeProductProperty
func GetGlobalAeopAeProductProperty() *GlobalAeopAeProductProperty {
	return poolGlobalAeopAeProductProperty.Get().(*GlobalAeopAeProductProperty)
}

// ReleaseGlobalAeopAeProductProperty 释放GlobalAeopAeProductProperty
func ReleaseGlobalAeopAeProductProperty(v *GlobalAeopAeProductProperty) {
	v.AttrName = ""
	v.AttrValue = ""
	v.AttrValueUnit = ""
	v.AttrNameId = 0
	v.AttrValueId = 0
	poolGlobalAeopAeProductProperty.Put(v)
}
