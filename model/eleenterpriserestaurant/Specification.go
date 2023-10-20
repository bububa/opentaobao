package eleenterpriserestaurant

import (
	"sync"
)

// Specification 结构体
type Specification struct {
	// 规格说明
	Values []string `json:"values,omitempty" xml:"values>string,omitempty"`
	// 特别说明
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

var poolSpecification = sync.Pool{
	New: func() any {
		return new(Specification)
	},
}

// GetSpecification() 从对象池中获取Specification
func GetSpecification() *Specification {
	return poolSpecification.Get().(*Specification)
}

// ReleaseSpecification 释放Specification
func ReleaseSpecification(v *Specification) {
	v.Values = v.Values[:0]
	v.Name = ""
	poolSpecification.Put(v)
}
