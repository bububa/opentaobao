package feedflow

import (
	"sync"
)

// OptionDto 结构体
type OptionDto struct {
	// 标签值
	OptionValue string `json:"option_value,omitempty" xml:"option_value,omitempty"`
	// 选项名称
	OptionName string `json:"option_name,omitempty" xml:"option_name,omitempty"`
	// 选项描述
	OptionDesc string `json:"option_desc,omitempty" xml:"option_desc,omitempty"`
}

var poolOptionDto = sync.Pool{
	New: func() any {
		return new(OptionDto)
	},
}

// GetOptionDto() 从对象池中获取OptionDto
func GetOptionDto() *OptionDto {
	return poolOptionDto.Get().(*OptionDto)
}

// ReleaseOptionDto 释放OptionDto
func ReleaseOptionDto(v *OptionDto) {
	v.OptionValue = ""
	v.OptionName = ""
	v.OptionDesc = ""
	poolOptionDto.Put(v)
}
