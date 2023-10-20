package simba

import (
	"sync"
)

// TagOptions 结构体
type TagOptions struct {
	// 标签的名称,如男,女,10-20等
	TagName string `json:"tag_name,omitempty" xml:"tag_name,omitempty"`
	// 标签id
	TagId string `json:"tag_id,omitempty" xml:"tag_id,omitempty"`
	// 标签所属的分组id
	OptionGroupId int64 `json:"option_group_id,omitempty" xml:"option_group_id,omitempty"`
	// 标签所属的维度id
	DimId int64 `json:"dim_id,omitempty" xml:"dim_id,omitempty"`
}

var poolTagOptions = sync.Pool{
	New: func() any {
		return new(TagOptions)
	},
}

// GetTagOptions() 从对象池中获取TagOptions
func GetTagOptions() *TagOptions {
	return poolTagOptions.Get().(*TagOptions)
}

// ReleaseTagOptions 释放TagOptions
func ReleaseTagOptions(v *TagOptions) {
	v.TagName = ""
	v.TagId = ""
	v.OptionGroupId = 0
	v.DimId = 0
	poolTagOptions.Put(v)
}
