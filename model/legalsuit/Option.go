package legalsuit

import (
	"sync"
)

// Option 结构体
type Option struct {
	// 文件对象
	FileValues []FileDto `json:"file_values,omitempty" xml:"file_values>file_dto,omitempty"`
	// 扩展字段标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 扩展字段值
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// 类型 string
	Type string `json:"type,omitempty" xml:"type,omitempty"`
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
	v.FileValues = v.FileValues[:0]
	v.Title = ""
	v.Value = ""
	v.Type = ""
	poolOption.Put(v)
}
