package legalcase

import (
	"sync"
)

// Option 结构体
type Option struct {
	// 标题值
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 属性值
	Value string `json:"value,omitempty" xml:"value,omitempty"`
}

var poolOption = sync.Pool{
	New: func() any {
		return new(Option)
	},
}

// GetOption() 从对象池中获取Option
func GetOption() *Option {
	return poolOption.Get().(*Option)
}

// ReleaseOption 释放Option
func ReleaseOption(v *Option) {
	v.Title = ""
	v.Value = ""
	poolOption.Put(v)
}
