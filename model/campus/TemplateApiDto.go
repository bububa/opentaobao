package campus

import (
	"sync"
)

// TemplateApiDto 结构体
type TemplateApiDto struct {
	// 参数点集合
	PropertyList []PropertyApiDto `json:"property_list,omitempty" xml:"property_list>property_api_dto,omitempty"`
	// 模板编码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 模板名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 模板id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolTemplateApiDto = sync.Pool{
	New: func() any {
		return new(TemplateApiDto)
	},
}

// GetTemplateApiDto() 从对象池中获取TemplateApiDto
func GetTemplateApiDto() *TemplateApiDto {
	return poolTemplateApiDto.Get().(*TemplateApiDto)
}

// ReleaseTemplateApiDto 释放TemplateApiDto
func ReleaseTemplateApiDto(v *TemplateApiDto) {
	v.PropertyList = v.PropertyList[:0]
	v.Code = ""
	v.Name = ""
	v.Id = 0
	poolTemplateApiDto.Put(v)
}
