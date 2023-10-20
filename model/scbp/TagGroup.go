package scbp

import (
	"sync"
)

// TagGroup 结构体
type TagGroup struct {
	// 分组名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 关键词数
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
	// 分组ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolTagGroup = sync.Pool{
	New: func() any {
		return new(TagGroup)
	},
}

// GetTagGroup() 从对象池中获取TagGroup
func GetTagGroup() *TagGroup {
	return poolTagGroup.Get().(*TagGroup)
}

// ReleaseTagGroup 释放TagGroup
func ReleaseTagGroup(v *TagGroup) {
	v.Name = ""
	v.Count = 0
	v.Id = 0
	poolTagGroup.Put(v)
}
