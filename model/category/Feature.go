package category

import (
	"sync"
)

// Feature 结构体
type Feature struct {
	// 属性键
	AttrKey string `json:"attr_key,omitempty" xml:"attr_key,omitempty"`
	// 属性值
	AttrValue string `json:"attr_value,omitempty" xml:"attr_value,omitempty"`
}

var poolFeature = sync.Pool{
	New: func() any {
		return new(Feature)
	},
}

// GetFeature() 从对象池中获取Feature
func GetFeature() *Feature {
	return poolFeature.Get().(*Feature)
}

// ReleaseFeature 释放Feature
func ReleaseFeature(v *Feature) {
	v.AttrKey = ""
	v.AttrValue = ""
	poolFeature.Put(v)
}
