package simba

import (
	"sync"
)

// Integer 结构体
type Integer struct {
	// 操作项code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 操作项名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 悬浮提示
	Tips string `json:"tips,omitempty" xml:"tips,omitempty"`
	// 是否可操作
	Disabled string `json:"disabled,omitempty" xml:"disabled,omitempty"`
	// 扩展属性
	Properties string `json:"properties,omitempty" xml:"properties,omitempty"`
}

var poolInteger = sync.Pool{
	New: func() any {
		return new(Integer)
	},
}

// GetInteger() 从对象池中获取Integer
func GetInteger() *Integer {
	return poolInteger.Get().(*Integer)
}

// ReleaseInteger 释放Integer
func ReleaseInteger(v *Integer) {
	v.Code = ""
	v.Name = ""
	v.Tips = ""
	v.Disabled = ""
	v.Properties = ""
	poolInteger.Put(v)
}
