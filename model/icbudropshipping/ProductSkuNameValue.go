package icbudropshipping

import (
	"sync"
)

// ProductSkuNameValue 结构体
type ProductSkuNameValue struct {
	// Attributes name
	AttrNameDesc string `json:"attr_name_desc,omitempty" xml:"attr_name_desc,omitempty"`
	// Attributes value
	AttrValueDesc string `json:"attr_value_desc,omitempty" xml:"attr_value_desc,omitempty"`
	// Attributes url
	AttrValueImage string `json:"attr_value_image,omitempty" xml:"attr_value_image,omitempty"`
	// Attributes name id
	AttrNameId int64 `json:"attr_name_id,omitempty" xml:"attr_name_id,omitempty"`
	// Attributes value id
	AttrValueId int64 `json:"attr_value_id,omitempty" xml:"attr_value_id,omitempty"`
}

var poolProductSkuNameValue = sync.Pool{
	New: func() any {
		return new(ProductSkuNameValue)
	},
}

// GetProductSkuNameValue() 从对象池中获取ProductSkuNameValue
func GetProductSkuNameValue() *ProductSkuNameValue {
	return poolProductSkuNameValue.Get().(*ProductSkuNameValue)
}

// ReleaseProductSkuNameValue 释放ProductSkuNameValue
func ReleaseProductSkuNameValue(v *ProductSkuNameValue) {
	v.AttrNameDesc = ""
	v.AttrValueDesc = ""
	v.AttrValueImage = ""
	v.AttrNameId = 0
	v.AttrValueId = 0
	poolProductSkuNameValue.Put(v)
}
