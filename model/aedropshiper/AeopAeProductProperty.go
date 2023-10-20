package aedropshiper

import (
	"sync"
)

// AeopAeProductProperty 结构体
type AeopAeProductProperty struct {
	// 属性单位
	AttrValueUnit string `json:"attr_value_unit,omitempty" xml:"attr_value_unit,omitempty"`
	// 区间属性开始值
	AttrValueStart string `json:"attr_value_start,omitempty" xml:"attr_value_start,omitempty"`
	// 区间属性结束值
	AttrValueEnd string `json:"attr_value_end,omitempty" xml:"attr_value_end,omitempty"`
	// 属性值
	AttrValue string `json:"attr_value,omitempty" xml:"attr_value,omitempty"`
	// 属性名字
	AttrName string `json:"attr_name,omitempty" xml:"attr_name,omitempty"`
	// 属性值ID
	AttrValueId int64 `json:"attr_value_id,omitempty" xml:"attr_value_id,omitempty"`
	// 属性Id
	AttrNameId int64 `json:"attr_name_id,omitempty" xml:"attr_name_id,omitempty"`
}

var poolAeopAeProductProperty = sync.Pool{
	New: func() any {
		return new(AeopAeProductProperty)
	},
}

// GetAeopAeProductProperty() 从对象池中获取AeopAeProductProperty
func GetAeopAeProductProperty() *AeopAeProductProperty {
	return poolAeopAeProductProperty.Get().(*AeopAeProductProperty)
}

// ReleaseAeopAeProductProperty 释放AeopAeProductProperty
func ReleaseAeopAeProductProperty(v *AeopAeProductProperty) {
	v.AttrValueUnit = ""
	v.AttrValueStart = ""
	v.AttrValueEnd = ""
	v.AttrValue = ""
	v.AttrName = ""
	v.AttrValueId = 0
	v.AttrNameId = 0
	poolAeopAeProductProperty.Put(v)
}
