package singletreasure

import (
	"sync"
)

// ActivityNameInfoDto 结构体
type ActivityNameInfoDto struct {
	// 活动名称的列表
	NameList []string `json:"name_list,omitempty" xml:"name_list>string,omitempty"`
	// 类目
	Category string `json:"category,omitempty" xml:"category,omitempty"`
}

var poolActivityNameInfoDto = sync.Pool{
	New: func() any {
		return new(ActivityNameInfoDto)
	},
}

// GetActivityNameInfoDto() 从对象池中获取ActivityNameInfoDto
func GetActivityNameInfoDto() *ActivityNameInfoDto {
	return poolActivityNameInfoDto.Get().(*ActivityNameInfoDto)
}

// ReleaseActivityNameInfoDto 释放ActivityNameInfoDto
func ReleaseActivityNameInfoDto(v *ActivityNameInfoDto) {
	v.NameList = v.NameList[:0]
	v.Category = ""
	poolActivityNameInfoDto.Put(v)
}
