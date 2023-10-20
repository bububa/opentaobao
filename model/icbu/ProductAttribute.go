package icbu

import (
	"sync"
)

// ProductAttribute 结构体
type ProductAttribute struct {
	// 属性名称
	AttributeName string `json:"attribute_name,omitempty" xml:"attribute_name,omitempty"`
	// 属性值名称
	ValueName string `json:"value_name,omitempty" xml:"value_name,omitempty"`
	// 作为sku属性值时，自定义属性值名称
	SkuCustomValueName string `json:"sku_custom_value_name,omitempty" xml:"sku_custom_value_name,omitempty"`
	// 作为sku属性值时，用图形来展示；必须是alibaba图片中心的图片URL，请使用API alibaba.icbu.photobank.upload 上传图片
	SkuCustomImageUrl string `json:"sku_custom_image_url,omitempty" xml:"sku_custom_image_url,omitempty"`
	// 属性ID
	AttributeId int64 `json:"attribute_id,omitempty" xml:"attribute_id,omitempty"`
	// 属性值ID
	ValueId int64 `json:"value_id,omitempty" xml:"value_id,omitempty"`
}

var poolProductAttribute = sync.Pool{
	New: func() any {
		return new(ProductAttribute)
	},
}

// GetProductAttribute() 从对象池中获取ProductAttribute
func GetProductAttribute() *ProductAttribute {
	return poolProductAttribute.Get().(*ProductAttribute)
}

// ReleaseProductAttribute 释放ProductAttribute
func ReleaseProductAttribute(v *ProductAttribute) {
	v.AttributeName = ""
	v.ValueName = ""
	v.SkuCustomValueName = ""
	v.SkuCustomImageUrl = ""
	v.AttributeId = 0
	v.ValueId = 0
	poolProductAttribute.Put(v)
}
