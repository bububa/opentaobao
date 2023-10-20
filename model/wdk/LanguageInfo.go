package wdk

import (
	"sync"
)

// LanguageInfo 结构体
type LanguageInfo struct {
	// 语言
	Language string `json:"language,omitempty" xml:"language,omitempty"`
	// 程度
	Degree string `json:"degree,omitempty" xml:"degree,omitempty"`
}

var poolLanguageInfo = sync.Pool{
	New: func() any {
		return new(LanguageInfo)
	},
}

// GetLanguageInfo() 从对象池中获取LanguageInfo
func GetLanguageInfo() *LanguageInfo {
	return poolLanguageInfo.Get().(*LanguageInfo)
}

// ReleaseLanguageInfo 释放LanguageInfo
func ReleaseLanguageInfo(v *LanguageInfo) {
	v.Language = ""
	v.Degree = ""
	poolLanguageInfo.Put(v)
}
