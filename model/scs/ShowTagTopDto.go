package scs

import (
	"sync"
)

// ShowTagTopDto 结构体
type ShowTagTopDto struct {
	// color
	Color string `json:"color,omitempty" xml:"color,omitempty"`
	// showTag
	ShowTag string `json:"show_tag,omitempty" xml:"show_tag,omitempty"`
	// tips
	Tips string `json:"tips,omitempty" xml:"tips,omitempty"`
	// tagId
	TagId int64 `json:"tag_id,omitempty" xml:"tag_id,omitempty"`
	// order
	Order int64 `json:"order,omitempty" xml:"order,omitempty"`
	// skip
	Skip bool `json:"skip,omitempty" xml:"skip,omitempty"`
}

var poolShowTagTopDto = sync.Pool{
	New: func() any {
		return new(ShowTagTopDto)
	},
}

// GetShowTagTopDto() 从对象池中获取ShowTagTopDto
func GetShowTagTopDto() *ShowTagTopDto {
	return poolShowTagTopDto.Get().(*ShowTagTopDto)
}

// ReleaseShowTagTopDto 释放ShowTagTopDto
func ReleaseShowTagTopDto(v *ShowTagTopDto) {
	v.Color = ""
	v.ShowTag = ""
	v.Tips = ""
	v.TagId = 0
	v.Order = 0
	v.Skip = false
	poolShowTagTopDto.Put(v)
}
