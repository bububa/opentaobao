package xiamicontent

import (
	"sync"
)

// TagCatLink 结构体
type TagCatLink struct {
	// 直属类目code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 直属类目中文名
	NameCn string `json:"name_cn,omitempty" xml:"name_cn,omitempty"`
	// 直属类目描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 直属类目英文名
	NameEn string `json:"name_en,omitempty" xml:"name_en,omitempty"`
	// tag根类目
	Parent *TagCatLink `json:"parent,omitempty" xml:"parent,omitempty"`
	// 直属类目根类目，如果为0则该类目为根类目
	Pid int64 `json:"pid,omitempty" xml:"pid,omitempty"`
	// 直属类目id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolTagCatLink = sync.Pool{
	New: func() any {
		return new(TagCatLink)
	},
}

// GetTagCatLink() 从对象池中获取TagCatLink
func GetTagCatLink() *TagCatLink {
	return poolTagCatLink.Get().(*TagCatLink)
}

// ReleaseTagCatLink 释放TagCatLink
func ReleaseTagCatLink(v *TagCatLink) {
	v.Code = ""
	v.NameCn = ""
	v.Description = ""
	v.NameEn = ""
	v.Parent = nil
	v.Pid = 0
	v.Id = 0
	poolTagCatLink.Put(v)
}
