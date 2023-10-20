package drug

import (
	"sync"
)

// Tags 结构体
type Tags struct {
	// type
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// picPath
	PicPath string `json:"pic_path,omitempty" xml:"pic_path,omitempty"`
	// desc
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// shortDesc
	ShortDesc string `json:"short_desc,omitempty" xml:"short_desc,omitempty"`
	// manFanType
	ManFanType string `json:"man_fan_type,omitempty" xml:"man_fan_type,omitempty"`
}

var poolTags = sync.Pool{
	New: func() any {
		return new(Tags)
	},
}

// GetTags() 从对象池中获取Tags
func GetTags() *Tags {
	return poolTags.Get().(*Tags)
}

// ReleaseTags 释放Tags
func ReleaseTags(v *Tags) {
	v.Type = ""
	v.PicPath = ""
	v.Desc = ""
	v.ShortDesc = ""
	v.ManFanType = ""
	poolTags.Put(v)
}
