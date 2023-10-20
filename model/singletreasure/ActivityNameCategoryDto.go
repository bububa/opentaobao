package singletreasure

import (
	"sync"
)

// ActivityNameCategoryDto 结构体
type ActivityNameCategoryDto struct {
	// 名称列表
	List []ActivityNameInfoDto `json:"list,omitempty" xml:"list>activity_name_info_dto,omitempty"`
	// 活动描述   1:日常活动  2:官方活动
	Text string `json:"text,omitempty" xml:"text,omitempty"`
	// 活动value
	Value int64 `json:"value,omitempty" xml:"value,omitempty"`
}

var poolActivityNameCategoryDto = sync.Pool{
	New: func() any {
		return new(ActivityNameCategoryDto)
	},
}

// GetActivityNameCategoryDto() 从对象池中获取ActivityNameCategoryDto
func GetActivityNameCategoryDto() *ActivityNameCategoryDto {
	return poolActivityNameCategoryDto.Get().(*ActivityNameCategoryDto)
}

// ReleaseActivityNameCategoryDto 释放ActivityNameCategoryDto
func ReleaseActivityNameCategoryDto(v *ActivityNameCategoryDto) {
	v.List = v.List[:0]
	v.Text = ""
	v.Value = 0
	poolActivityNameCategoryDto.Put(v)
}
