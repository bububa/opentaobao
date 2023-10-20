package legalcase

import (
	"sync"
)

// Content 结构体
type Content struct {
	// 二级值集
	Children []Children `json:"children,omitempty" xml:"children>children,omitempty"`
	// code值
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// 文本值
	Text string `json:"text,omitempty" xml:"text,omitempty"`
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
	v.Children = v.Children[:0]
	v.Value = ""
	v.Text = ""
	poolContent.Put(v)
}
