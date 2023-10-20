package util

import (
	"sync"
)

// ShortUrlOption 结构体
type ShortUrlOption struct {
	// bizName
	BizName string `json:"biz_name,omitempty" xml:"biz_name,omitempty"`
}

var poolShortUrlOption = sync.Pool{
	New: func() any {
		return new(ShortUrlOption)
	},
}

// GetShortUrlOption() 从对象池中获取ShortUrlOption
func GetShortUrlOption() *ShortUrlOption {
	return poolShortUrlOption.Get().(*ShortUrlOption)
}

// ReleaseShortUrlOption 释放ShortUrlOption
func ReleaseShortUrlOption(v *ShortUrlOption) {
	v.BizName = ""
	poolShortUrlOption.Put(v)
}
