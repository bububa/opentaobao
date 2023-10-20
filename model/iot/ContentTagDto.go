package iot

import (
	"sync"
)

// ContentTagDto 结构体
type ContentTagDto struct {
	// 标签名称
	TagName string `json:"tag_name,omitempty" xml:"tag_name,omitempty"`
	// 标签id
	TagId int64 `json:"tag_id,omitempty" xml:"tag_id,omitempty"`
}

var poolContentTagDto = sync.Pool{
	New: func() any {
		return new(ContentTagDto)
	},
}

// GetContentTagDto() 从对象池中获取ContentTagDto
func GetContentTagDto() *ContentTagDto {
	return poolContentTagDto.Get().(*ContentTagDto)
}

// ReleaseContentTagDto 释放ContentTagDto
func ReleaseContentTagDto(v *ContentTagDto) {
	v.TagName = ""
	v.TagId = 0
	poolContentTagDto.Put(v)
}
