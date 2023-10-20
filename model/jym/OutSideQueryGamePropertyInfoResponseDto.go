package jym

import (
	"sync"
)

// OutSideQueryGamePropertyInfoResponseDto 结构体
type OutSideQueryGamePropertyInfoResponseDto struct {
	// 加密压缩的属性信息
	GamePropertiesCompressed string `json:"game_properties_compressed,omitempty" xml:"game_properties_compressed,omitempty"`
}

var poolOutSideQueryGamePropertyInfoResponseDto = sync.Pool{
	New: func() any {
		return new(OutSideQueryGamePropertyInfoResponseDto)
	},
}

// GetOutSideQueryGamePropertyInfoResponseDto() 从对象池中获取OutSideQueryGamePropertyInfoResponseDto
func GetOutSideQueryGamePropertyInfoResponseDto() *OutSideQueryGamePropertyInfoResponseDto {
	return poolOutSideQueryGamePropertyInfoResponseDto.Get().(*OutSideQueryGamePropertyInfoResponseDto)
}

// ReleaseOutSideQueryGamePropertyInfoResponseDto 释放OutSideQueryGamePropertyInfoResponseDto
func ReleaseOutSideQueryGamePropertyInfoResponseDto(v *OutSideQueryGamePropertyInfoResponseDto) {
	v.GamePropertiesCompressed = ""
	poolOutSideQueryGamePropertyInfoResponseDto.Put(v)
}
