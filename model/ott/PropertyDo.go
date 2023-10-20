package ott

import (
	"sync"
)

// PropertyDo 结构体
type PropertyDo struct {
	// 属性键值对
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 属性校验码
	VerifyCode string `json:"verify_code,omitempty" xml:"verify_code,omitempty"`
}

var poolPropertyDo = sync.Pool{
	New: func() any {
		return new(PropertyDo)
	},
}

// GetPropertyDo() 从对象池中获取PropertyDo
func GetPropertyDo() *PropertyDo {
	return poolPropertyDo.Get().(*PropertyDo)
}

// ReleasePropertyDo 释放PropertyDo
func ReleasePropertyDo(v *PropertyDo) {
	v.Data = ""
	v.VerifyCode = ""
	poolPropertyDo.Put(v)
}
