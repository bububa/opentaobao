package aedropshiper

import (
	"sync"
)

// AeopSkuProperty 结构体
type AeopSkuProperty struct {
	// SKU图片地址
	SkuImage string `json:"sku_image,omitempty" xml:"sku_image,omitempty"`
	// 属性值自定义名称
	PropertyValueDefinitionName string `json:"property_value_definition_name,omitempty" xml:"property_value_definition_name,omitempty"`
	// 属性值
	SkuPropertyValue string `json:"sku_property_value,omitempty" xml:"sku_property_value,omitempty"`
	// 属性名
	SkuPropertyName string `json:"sku_property_name,omitempty" xml:"sku_property_name,omitempty"`
	// SKU属性ID
	SkuPropertyId int64 `json:"sku_property_id,omitempty" xml:"sku_property_id,omitempty"`
	// SKU属性值ID
	PropertyValueId int64 `json:"property_value_id,omitempty" xml:"property_value_id,omitempty"`
	// 自定义id
	PropertyValueIdLong int64 `json:"property_value_id_long,omitempty" xml:"property_value_id_long,omitempty"`
}

var poolAeopSkuProperty = sync.Pool{
	New: func() any {
		return new(AeopSkuProperty)
	},
}

// GetAeopSkuProperty() 从对象池中获取AeopSkuProperty
func GetAeopSkuProperty() *AeopSkuProperty {
	return poolAeopSkuProperty.Get().(*AeopSkuProperty)
}

// ReleaseAeopSkuProperty 释放AeopSkuProperty
func ReleaseAeopSkuProperty(v *AeopSkuProperty) {
	v.SkuImage = ""
	v.PropertyValueDefinitionName = ""
	v.SkuPropertyValue = ""
	v.SkuPropertyName = ""
	v.SkuPropertyId = 0
	v.PropertyValueId = 0
	v.PropertyValueIdLong = 0
	poolAeopSkuProperty.Put(v)
}
