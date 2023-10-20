package jst

import (
	"sync"
)

// DigitalSmsTemplateContentDto 结构体
type DigitalSmsTemplateContentDto struct {
	// 文件名称
	FileName string `json:"file_name,omitempty" xml:"file_name,omitempty"`
	// 文件后缀名
	FileSuffix string `json:"file_suffix,omitempty" xml:"file_suffix,omitempty"`
	// 文件Base64转码后的字符串
	FileContents string `json:"file_contents,omitempty" xml:"file_contents,omitempty"`
	// 文件大小
	FileSize int64 `json:"file_size,omitempty" xml:"file_size,omitempty"`
}

var poolDigitalSmsTemplateContentDto = sync.Pool{
	New: func() any {
		return new(DigitalSmsTemplateContentDto)
	},
}

// GetDigitalSmsTemplateContentDto() 从对象池中获取DigitalSmsTemplateContentDto
func GetDigitalSmsTemplateContentDto() *DigitalSmsTemplateContentDto {
	return poolDigitalSmsTemplateContentDto.Get().(*DigitalSmsTemplateContentDto)
}

// ReleaseDigitalSmsTemplateContentDto 释放DigitalSmsTemplateContentDto
func ReleaseDigitalSmsTemplateContentDto(v *DigitalSmsTemplateContentDto) {
	v.FileName = ""
	v.FileSuffix = ""
	v.FileContents = ""
	v.FileSize = 0
	poolDigitalSmsTemplateContentDto.Put(v)
}
