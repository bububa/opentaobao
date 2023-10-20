package alihouse

import (
	"sync"
)

// TagsRequestDto 结构体
type TagsRequestDto struct {
	// 标列表
	AddTags []string `json:"add_tags,omitempty" xml:"add_tags>string,omitempty"`
	// 商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}

var poolTagsRequestDto = sync.Pool{
	New: func() any {
		return new(TagsRequestDto)
	},
}

// GetTagsRequestDto() 从对象池中获取TagsRequestDto
func GetTagsRequestDto() *TagsRequestDto {
	return poolTagsRequestDto.Get().(*TagsRequestDto)
}

// ReleaseTagsRequestDto 释放TagsRequestDto
func ReleaseTagsRequestDto(v *TagsRequestDto) {
	v.AddTags = v.AddTags[:0]
	v.ItemId = 0
	poolTagsRequestDto.Put(v)
}
