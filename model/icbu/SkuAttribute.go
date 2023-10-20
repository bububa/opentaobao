package icbu

import (
	"sync"
)

// SkuAttribute 结构体
type SkuAttribute struct {
	// 属性下的值
	Values []SkuAttributeValue `json:"values,omitempty" xml:"values>sku_attribute_value,omitempty"`
	// 属性名称
	AttributeName string `json:"attribute_name,omitempty" xml:"attribute_name,omitempty"`
	// 属性ID
	AttributeId int64 `json:"attribute_id,omitempty" xml:"attribute_id,omitempty"`
}

var poolSkuAttribute = sync.Pool{
	New: func() any {
		return new(SkuAttribute)
	},
}

// GetSkuAttribute() 从对象池中获取SkuAttribute
func GetSkuAttribute() *SkuAttribute {
	return poolSkuAttribute.Get().(*SkuAttribute)
}

// ReleaseSkuAttribute 释放SkuAttribute
func ReleaseSkuAttribute(v *SkuAttribute) {
	v.Values = v.Values[:0]
	v.AttributeName = ""
	v.AttributeId = 0
	poolSkuAttribute.Put(v)
}
