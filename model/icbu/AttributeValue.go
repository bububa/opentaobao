package icbu

import (
	"sync"
)

// AttributeValue 结构体
type AttributeValue struct {
	// 该属性值的子属性id
	ChildAttrs []string `json:"child_attrs,omitempty" xml:"child_attrs>string,omitempty"`
	// 英文名字
	EnName string `json:"en_name,omitempty" xml:"en_name,omitempty"`
	// 属性值id
	AttrValueId int64 `json:"attr_value_id,omitempty" xml:"attr_value_id,omitempty"`
	// 属性id
	AttrId int64 `json:"attr_id,omitempty" xml:"attr_id,omitempty"`
	// 所属类目id
	CatId int64 `json:"cat_id,omitempty" xml:"cat_id,omitempty"`
	// 是否SKU属性值
	SkuValue bool `json:"sku_value,omitempty" xml:"sku_value,omitempty"`
}

var poolAttributeValue = sync.Pool{
	New: func() any {
		return new(AttributeValue)
	},
}

// GetAttributeValue() 从对象池中获取AttributeValue
func GetAttributeValue() *AttributeValue {
	return poolAttributeValue.Get().(*AttributeValue)
}

// ReleaseAttributeValue 释放AttributeValue
func ReleaseAttributeValue(v *AttributeValue) {
	v.ChildAttrs = v.ChildAttrs[:0]
	v.EnName = ""
	v.AttrValueId = 0
	v.AttrId = 0
	v.CatId = 0
	v.SkuValue = false
	poolAttributeValue.Put(v)
}
