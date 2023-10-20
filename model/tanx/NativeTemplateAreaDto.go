package tanx

import (
	"sync"
)

// NativeTemplateAreaDto 结构体
type NativeTemplateAreaDto struct {
	// 广告区域ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 创意个数
	CreativeCount int64 `json:"creative_count,omitempty" xml:"creative_count,omitempty"`
	// 创意要求
	Creative *NativeTemplateCreativeDto `json:"creative,omitempty" xml:"creative,omitempty"`
}

var poolNativeTemplateAreaDto = sync.Pool{
	New: func() any {
		return new(NativeTemplateAreaDto)
	},
}

// GetNativeTemplateAreaDto() 从对象池中获取NativeTemplateAreaDto
func GetNativeTemplateAreaDto() *NativeTemplateAreaDto {
	return poolNativeTemplateAreaDto.Get().(*NativeTemplateAreaDto)
}

// ReleaseNativeTemplateAreaDto 释放NativeTemplateAreaDto
func ReleaseNativeTemplateAreaDto(v *NativeTemplateAreaDto) {
	v.Id = 0
	v.CreativeCount = 0
	v.Creative = nil
	poolNativeTemplateAreaDto.Put(v)
}
