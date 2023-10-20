package omniorder

import (
	"sync"
)

// Content 结构体
type Content struct {
	// 取件码
	Key string `json:"key,omitempty" xml:"key,omitempty"`
	// 取件码的值
	Value string `json:"value,omitempty" xml:"value,omitempty"`
}

var poolContent = sync.Pool{
	New: func() any {
		return new(Content)
	},
}

// GetContent() 从对象池中获取Content
func GetContent() *Content {
	return poolContent.Get().(*Content)
}

// ReleaseContent 释放Content
func ReleaseContent(v *Content) {
	v.Key = ""
	v.Value = ""
	poolContent.Put(v)
}
