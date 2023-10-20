package ott

import (
	"sync"
)

// AvailableLanguageDto 结构体
type AvailableLanguageDto struct {
	// availableLanguage
	AvailableLanguage string `json:"available_language,omitempty" xml:"available_language,omitempty"`
}

var poolAvailableLanguageDto = sync.Pool{
	New: func() any {
		return new(AvailableLanguageDto)
	},
}

// GetAvailableLanguageDto() 从对象池中获取AvailableLanguageDto
func GetAvailableLanguageDto() *AvailableLanguageDto {
	return poolAvailableLanguageDto.Get().(*AvailableLanguageDto)
}

// ReleaseAvailableLanguageDto 释放AvailableLanguageDto
func ReleaseAvailableLanguageDto(v *AvailableLanguageDto) {
	v.AvailableLanguage = ""
	poolAvailableLanguageDto.Put(v)
}
