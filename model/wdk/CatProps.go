package wdk

import (
	"sync"
)

// CatProps 结构体
type CatProps struct {
	// 类目属性名称
	PropertyText string `json:"property_text,omitempty" xml:"property_text,omitempty"`
	// 类目值名称
	ValueText string `json:"value_text,omitempty" xml:"value_text,omitempty"`
	// 类目属性ID
	PropertyId int64 `json:"property_id,omitempty" xml:"property_id,omitempty"`
	// 类目值ID
	ValueId int64 `json:"value_id,omitempty" xml:"value_id,omitempty"`
}

var poolCatProps = sync.Pool{
	New: func() any {
		return new(CatProps)
	},
}

// GetCatProps() 从对象池中获取CatProps
func GetCatProps() *CatProps {
	return poolCatProps.Get().(*CatProps)
}

// ReleaseCatProps 释放CatProps
func ReleaseCatProps(v *CatProps) {
	v.PropertyText = ""
	v.ValueText = ""
	v.PropertyId = 0
	v.ValueId = 0
	poolCatProps.Put(v)
}
