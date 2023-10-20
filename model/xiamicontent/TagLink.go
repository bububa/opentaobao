package xiamicontent

import (
	"sync"
)

// TagLink 结构体
type TagLink struct {
	// 标签code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 标签名
	NameCn string `json:"name_cn,omitempty" xml:"name_cn,omitempty"`
	// tag描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// tag英文名
	NameEn string `json:"name_en,omitempty" xml:"name_en,omitempty"`
	// 标签id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 父标签id
	Pid int64 `json:"pid,omitempty" xml:"pid,omitempty"`
	// tag直属类目
	Parent *TagCatLink `json:"parent,omitempty" xml:"parent,omitempty"`
}

var poolTagLink = sync.Pool{
	New: func() any {
		return new(TagLink)
	},
}

// GetTagLink() 从对象池中获取TagLink
func GetTagLink() *TagLink {
	return poolTagLink.Get().(*TagLink)
}

// ReleaseTagLink 释放TagLink
func ReleaseTagLink(v *TagLink) {
	v.Code = ""
	v.NameCn = ""
	v.Description = ""
	v.NameEn = ""
	v.Id = 0
	v.Pid = 0
	v.Parent = nil
	poolTagLink.Put(v)
}
