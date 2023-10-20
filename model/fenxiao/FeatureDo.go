package fenxiao

import (
	"sync"
)

// FeatureDo 结构体
type FeatureDo struct {
	// 属性键
	AttrKey string `json:"attr_key,omitempty" xml:"attr_key,omitempty"`
	// 属性值
	AttrValue string `json:"attr_value,omitempty" xml:"attr_value,omitempty"`
}

var poolFeatureDo = sync.Pool{
	New: func() any {
		return new(FeatureDo)
	},
}

// GetFeatureDo() 从对象池中获取FeatureDo
func GetFeatureDo() *FeatureDo {
	return poolFeatureDo.Get().(*FeatureDo)
}

// ReleaseFeatureDo 释放FeatureDo
func ReleaseFeatureDo(v *FeatureDo) {
	v.AttrKey = ""
	v.AttrValue = ""
	poolFeatureDo.Put(v)
}
