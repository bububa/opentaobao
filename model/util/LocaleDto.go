package util

import (
	"sync"
)

// LocaleDto 结构体
type LocaleDto struct {
	// Locale编码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// Locale名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

var poolLocaleDto = sync.Pool{
	New: func() any {
		return new(LocaleDto)
	},
}

// GetLocaleDto() 从对象池中获取LocaleDto
func GetLocaleDto() *LocaleDto {
	return poolLocaleDto.Get().(*LocaleDto)
}

// ReleaseLocaleDto 释放LocaleDto
func ReleaseLocaleDto(v *LocaleDto) {
	v.Code = ""
	v.Name = ""
	poolLocaleDto.Put(v)
}
