package eleenterpriserestaurant

import (
	"sync"
)

// Attr 结构体
type Attr struct {
	// 值
	Values []string `json:"values,omitempty" xml:"values>string,omitempty"`
	// 姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

var poolAttr = sync.Pool{
	New: func() any {
		return new(Attr)
	},
}

// GetAttr() 从对象池中获取Attr
func GetAttr() *Attr {
	return poolAttr.Get().(*Attr)
}

// ReleaseAttr 释放Attr
func ReleaseAttr(v *Attr) {
	v.Values = v.Values[:0]
	v.Name = ""
	poolAttr.Put(v)
}
