package icbu

import (
	"sync"
)

// CustomContent 结构体
type CustomContent struct {
	// 定制类型，只允许填写英文字符
	CustomType string `json:"custom_type,omitempty" xml:"custom_type,omitempty"`
	// 最小起订量
	MinOrderQuantity int64 `json:"min_order_quantity,omitempty" xml:"min_order_quantity,omitempty"`
}

var poolCustomContent = sync.Pool{
	New: func() any {
		return new(CustomContent)
	},
}

// GetCustomContent() 从对象池中获取CustomContent
func GetCustomContent() *CustomContent {
	return poolCustomContent.Get().(*CustomContent)
}

// ReleaseCustomContent 释放CustomContent
func ReleaseCustomContent(v *CustomContent) {
	v.CustomType = ""
	v.MinOrderQuantity = 0
	poolCustomContent.Put(v)
}
