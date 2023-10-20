package idleisv

import (
	"sync"
)

// IdleItemApiPvPairDo 结构体
type IdleItemApiPvPairDo struct {
	// 属性名文本(长度&lt;=30)
	PropertyText string `json:"property_text,omitempty" xml:"property_text,omitempty"`
	// 属性值文本(长度&lt;=30)
	ValueText string `json:"value_text,omitempty" xml:"value_text,omitempty"`
}

var poolIdleItemApiPvPairDo = sync.Pool{
	New: func() any {
		return new(IdleItemApiPvPairDo)
	},
}

// GetIdleItemApiPvPairDo() 从对象池中获取IdleItemApiPvPairDo
func GetIdleItemApiPvPairDo() *IdleItemApiPvPairDo {
	return poolIdleItemApiPvPairDo.Get().(*IdleItemApiPvPairDo)
}

// ReleaseIdleItemApiPvPairDo 释放IdleItemApiPvPairDo
func ReleaseIdleItemApiPvPairDo(v *IdleItemApiPvPairDo) {
	v.PropertyText = ""
	v.ValueText = ""
	poolIdleItemApiPvPairDo.Put(v)
}
