package legalcase

import (
	"sync"
)

// Children 结构体
type Children struct {
	// 文本值
	Text string `json:"text,omitempty" xml:"text,omitempty"`
	// code值
	Value string `json:"value,omitempty" xml:"value,omitempty"`
}

var poolChildren = sync.Pool{
	New: func() any {
		return new(Children)
	},
}

// GetChildren() 从对象池中获取Children
func GetChildren() *Children {
	return poolChildren.Get().(*Children)
}

// ReleaseChildren 释放Children
func ReleaseChildren(v *Children) {
	v.Text = ""
	v.Value = ""
	poolChildren.Put(v)
}
