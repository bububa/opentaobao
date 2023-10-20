package aedropshiper

import (
	"sync"
)

// AeSkuPropertyDto 结构体
type AeSkuPropertyDto struct {
	// Attribute name
	SkuPropertyName string `json:"sku_property_name,omitempty" xml:"sku_property_name,omitempty"`
	// Attribute value
	SkuPropertyValue string `json:"sku_property_value,omitempty" xml:"sku_property_value,omitempty"`
	// Custom name
	PropertyValueDefinitionName string `json:"property_value_definition_name,omitempty" xml:"property_value_definition_name,omitempty"`
	// SKU pictures
	SkuImage string `json:"sku_image,omitempty" xml:"sku_image,omitempty"`
	// Attribute ID
	SkuPropertyId int64 `json:"sku_property_id,omitempty" xml:"sku_property_id,omitempty"`
	// Custom id
	PropertyValueId int64 `json:"property_value_id,omitempty" xml:"property_value_id,omitempty"`
}

var poolAeSkuPropertyDto = sync.Pool{
	New: func() any {
		return new(AeSkuPropertyDto)
	},
}

// GetAeSkuPropertyDto() 从对象池中获取AeSkuPropertyDto
func GetAeSkuPropertyDto() *AeSkuPropertyDto {
	return poolAeSkuPropertyDto.Get().(*AeSkuPropertyDto)
}

// ReleaseAeSkuPropertyDto 释放AeSkuPropertyDto
func ReleaseAeSkuPropertyDto(v *AeSkuPropertyDto) {
	v.SkuPropertyName = ""
	v.SkuPropertyValue = ""
	v.PropertyValueDefinitionName = ""
	v.SkuImage = ""
	v.SkuPropertyId = 0
	v.PropertyValueId = 0
	poolAeSkuPropertyDto.Put(v)
}
