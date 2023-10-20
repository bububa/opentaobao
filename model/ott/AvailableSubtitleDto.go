package ott

import (
	"sync"
)

// AvailableSubtitleDto 结构体
type AvailableSubtitleDto struct {
	// availableSubtitle
	AvailableSubtitle string `json:"available_subtitle,omitempty" xml:"available_subtitle,omitempty"`
}

var poolAvailableSubtitleDto = sync.Pool{
	New: func() any {
		return new(AvailableSubtitleDto)
	},
}

// GetAvailableSubtitleDto() 从对象池中获取AvailableSubtitleDto
func GetAvailableSubtitleDto() *AvailableSubtitleDto {
	return poolAvailableSubtitleDto.Get().(*AvailableSubtitleDto)
}

// ReleaseAvailableSubtitleDto 释放AvailableSubtitleDto
func ReleaseAvailableSubtitleDto(v *AvailableSubtitleDto) {
	v.AvailableSubtitle = ""
	poolAvailableSubtitleDto.Put(v)
}
