package scbp

import (
	"sync"
)

// TagDefineDto 结构体
type TagDefineDto struct {
	// 标签值
	OptionValue string `json:"option_value,omitempty" xml:"option_value,omitempty"`
	// 标签名（标签描述为空时，取标签名）
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 标签描述（标签名为空时，取标签描述）
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 层级（0,1,2）
	Level int64 `json:"level,omitempty" xml:"level,omitempty"`
}

var poolTagDefineDto = sync.Pool{
	New: func() any {
		return new(TagDefineDto)
	},
}

// GetTagDefineDto() 从对象池中获取TagDefineDto
func GetTagDefineDto() *TagDefineDto {
	return poolTagDefineDto.Get().(*TagDefineDto)
}

// ReleaseTagDefineDto 释放TagDefineDto
func ReleaseTagDefineDto(v *TagDefineDto) {
	v.OptionValue = ""
	v.Name = ""
	v.Desc = ""
	v.Level = 0
	poolTagDefineDto.Put(v)
}
