package aedropshiper

// AeopSkuProperty 结构体
type AeopSkuProperty struct {
	// SKU属性ID
	SkuPropertyId int64 `json:"sku_property_id,omitempty" xml:"sku_property_id,omitempty"`
	// SKU图片地址
	SkuImage string `json:"sku_image,omitempty" xml:"sku_image,omitempty"`
	// SKU属性值ID
	PropertyValueId int64 `json:"property_value_id,omitempty" xml:"property_value_id,omitempty"`
	// 属性值自定义名称
	PropertyValueDefinitionName string `json:"property_value_definition_name,omitempty" xml:"property_value_definition_name,omitempty"`
	// 自定义id
	PropertyValueIdLong int64 `json:"property_value_id_long,omitempty" xml:"property_value_id_long,omitempty"`
	// 属性值
	SkuPropertyValue string `json:"sku_property_value,omitempty" xml:"sku_property_value,omitempty"`
	// 属性名
	SkuPropertyName string `json:"sku_property_name,omitempty" xml:"sku_property_name,omitempty"`
}
