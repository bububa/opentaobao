package scbp

import (
	"sync"
)

// ErrorKeyword 结构体
type ErrorKeyword struct {
	// 关键词
	Keyword string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// 类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 重复关键词
	RepeatKeyword string `json:"repeat_keyword,omitempty" xml:"repeat_keyword,omitempty"`
}

var poolErrorKeyword = sync.Pool{
	New: func() any {
		return new(ErrorKeyword)
	},
}

// GetErrorKeyword() 从对象池中获取ErrorKeyword
func GetErrorKeyword() *ErrorKeyword {
	return poolErrorKeyword.Get().(*ErrorKeyword)
}

// ReleaseErrorKeyword 释放ErrorKeyword
func ReleaseErrorKeyword(v *ErrorKeyword) {
	v.Keyword = ""
	v.Type = ""
	v.RepeatKeyword = ""
	poolErrorKeyword.Put(v)
}
