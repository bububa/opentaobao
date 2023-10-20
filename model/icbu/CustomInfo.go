package icbu

import (
	"sync"
)

// CustomInfo 结构体
type CustomInfo struct {
	// 定制内容
	CustomContents []CustomContent `json:"custom_contents,omitempty" xml:"custom_contents>custom_content,omitempty"`
}

var poolCustomInfo = sync.Pool{
	New: func() any {
		return new(CustomInfo)
	},
}

// GetCustomInfo() 从对象池中获取CustomInfo
func GetCustomInfo() *CustomInfo {
	return poolCustomInfo.Get().(*CustomInfo)
}

// ReleaseCustomInfo 释放CustomInfo
func ReleaseCustomInfo(v *CustomInfo) {
	v.CustomContents = v.CustomContents[:0]
	poolCustomInfo.Put(v)
}
