package aesolution

import (
	"sync"
)

// GlobalAeopSkuProperty 结构体
type GlobalAeopSkuProperty struct {
	// Customized attribute value name
	PropertyValueDefinitionName string `json:"property_value_definition_name,omitempty" xml:"property_value_definition_name,omitempty"`
	// Self-defined image url for this sku.
	SkuImage string `json:"sku_image,omitempty" xml:"sku_image,omitempty"`
	// SKU attribute value id
	PropertyValueId int64 `json:"property_value_id,omitempty" xml:"property_value_id,omitempty"`
	// SKU attribute name id
	SkuPropertyId int64 `json:"sku_property_id,omitempty" xml:"sku_property_id,omitempty"`
}

var poolGlobalAeopSkuProperty = sync.Pool{
	New: func() any {
		return new(GlobalAeopSkuProperty)
	},
}

// GetGlobalAeopSkuProperty() 从对象池中获取GlobalAeopSkuProperty
func GetGlobalAeopSkuProperty() *GlobalAeopSkuProperty {
	return poolGlobalAeopSkuProperty.Get().(*GlobalAeopSkuProperty)
}

// ReleaseGlobalAeopSkuProperty 释放GlobalAeopSkuProperty
func ReleaseGlobalAeopSkuProperty(v *GlobalAeopSkuProperty) {
	v.PropertyValueDefinitionName = ""
	v.SkuImage = ""
	v.PropertyValueId = 0
	v.SkuPropertyId = 0
	poolGlobalAeopSkuProperty.Put(v)
}
