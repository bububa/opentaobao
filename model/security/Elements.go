package security

import (
	"sync"
)

// Elements 结构体
type Elements struct {
	// 材料值类型
	ValueMeta string `json:"value_meta,omitempty" xml:"value_meta,omitempty"`
	// url
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// 材料名字
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

var poolElements = sync.Pool{
	New: func() any {
		return new(Elements)
	},
}

// GetElements() 从对象池中获取Elements
func GetElements() *Elements {
	return poolElements.Get().(*Elements)
}

// ReleaseElements 释放Elements
func ReleaseElements(v *Elements) {
	v.ValueMeta = ""
	v.Value = ""
	v.Name = ""
	poolElements.Put(v)
}
