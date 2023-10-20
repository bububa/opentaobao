package eleenterpriserestaurant

import (
	"sync"
)

// Spec 结构体
type Spec struct {
	// 姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 值
	Value string `json:"value,omitempty" xml:"value,omitempty"`
}

var poolSpec = sync.Pool{
	New: func() any {
		return new(Spec)
	},
}

// GetSpec() 从对象池中获取Spec
func GetSpec() *Spec {
	return poolSpec.Get().(*Spec)
}

// ReleaseSpec 释放Spec
func ReleaseSpec(v *Spec) {
	v.Name = ""
	v.Value = ""
	poolSpec.Put(v)
}
