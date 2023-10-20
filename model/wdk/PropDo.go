package wdk

import (
	"sync"
)

// PropDo 结构体
type PropDo struct {
	// 行业属性
	Key string `json:"key,omitempty" xml:"key,omitempty"`
	// 类目属性值
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// 是否移除该属性
	RemoveOpt bool `json:"remove_opt,omitempty" xml:"remove_opt,omitempty"`
}

var poolPropDo = sync.Pool{
	New: func() any {
		return new(PropDo)
	},
}

// GetPropDo() 从对象池中获取PropDo
func GetPropDo() *PropDo {
	return poolPropDo.Get().(*PropDo)
}

// ReleasePropDo 释放PropDo
func ReleasePropDo(v *PropDo) {
	v.Key = ""
	v.Value = ""
	v.RemoveOpt = false
	poolPropDo.Put(v)
}
