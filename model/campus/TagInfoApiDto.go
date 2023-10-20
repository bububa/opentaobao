package campus

import (
	"sync"
)

// TagInfoApiDto 结构体
type TagInfoApiDto struct {
	// 标签名称
	TagName string `json:"tag_name,omitempty" xml:"tag_name,omitempty"`
	// 标签描述
	TagDesc string `json:"tag_desc,omitempty" xml:"tag_desc,omitempty"`
	// 标签类型
	TagTypeName string `json:"tag_type_name,omitempty" xml:"tag_type_name,omitempty"`
	// 是否系统标签
	SystemTag bool `json:"system_tag,omitempty" xml:"system_tag,omitempty"`
}

var poolTagInfoApiDto = sync.Pool{
	New: func() any {
		return new(TagInfoApiDto)
	},
}

// GetTagInfoApiDto() 从对象池中获取TagInfoApiDto
func GetTagInfoApiDto() *TagInfoApiDto {
	return poolTagInfoApiDto.Get().(*TagInfoApiDto)
}

// ReleaseTagInfoApiDto 释放TagInfoApiDto
func ReleaseTagInfoApiDto(v *TagInfoApiDto) {
	v.TagName = ""
	v.TagDesc = ""
	v.TagTypeName = ""
	v.SystemTag = false
	poolTagInfoApiDto.Put(v)
}
