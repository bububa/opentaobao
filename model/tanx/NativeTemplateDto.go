package tanx

import (
	"sync"
)

// NativeTemplateDto 结构体
type NativeTemplateDto struct {
	// 区域列表
	Areas []NativeTemplateAreaDto `json:"areas,omitempty" xml:"areas>native_template_area_dto,omitempty"`
	// 样式预览图片url
	Preview string `json:"preview,omitempty" xml:"preview,omitempty"`
	// 说明
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 模板支持的广告位尺寸
	Size string `json:"size,omitempty" xml:"size,omitempty"`
	// 模板ID
	TmplId int64 `json:"tmpl_id,omitempty" xml:"tmpl_id,omitempty"`
}

var poolNativeTemplateDto = sync.Pool{
	New: func() any {
		return new(NativeTemplateDto)
	},
}

// GetNativeTemplateDto() 从对象池中获取NativeTemplateDto
func GetNativeTemplateDto() *NativeTemplateDto {
	return poolNativeTemplateDto.Get().(*NativeTemplateDto)
}

// ReleaseNativeTemplateDto 释放NativeTemplateDto
func ReleaseNativeTemplateDto(v *NativeTemplateDto) {
	v.Areas = v.Areas[:0]
	v.Preview = ""
	v.Description = ""
	v.Size = ""
	v.TmplId = 0
	poolNativeTemplateDto.Put(v)
}
