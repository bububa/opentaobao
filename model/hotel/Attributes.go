package hotel

import (
	"sync"
)

// Attributes 结构体
type Attributes struct {
	// key
	Key string `json:"key,omitempty" xml:"key,omitempty"`
}

var poolAttributes = sync.Pool{
	New: func() any {
		return new(Attributes)
	},
}

// GetAttributes() 从对象池中获取Attributes
func GetAttributes() *Attributes {
	return poolAttributes.Get().(*Attributes)
}

// ReleaseAttributes 释放Attributes
func ReleaseAttributes(v *Attributes) {
	v.Key = ""
	poolAttributes.Put(v)
}
