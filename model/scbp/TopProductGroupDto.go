package scbp

import (
	"sync"
)

// TopProductGroupDto 结构体
type TopProductGroupDto struct {
	// 产品分组标识
	GroupId string `json:"group_id,omitempty" xml:"group_id,omitempty"`
	// 产品分组名称
	GroupName string `json:"group_name,omitempty" xml:"group_name,omitempty"`
	// 是否是叶子分组，即没有子分组
	Leaf bool `json:"leaf,omitempty" xml:"leaf,omitempty"`
}

var poolTopProductGroupDto = sync.Pool{
	New: func() any {
		return new(TopProductGroupDto)
	},
}

// GetTopProductGroupDto() 从对象池中获取TopProductGroupDto
func GetTopProductGroupDto() *TopProductGroupDto {
	return poolTopProductGroupDto.Get().(*TopProductGroupDto)
}

// ReleaseTopProductGroupDto 释放TopProductGroupDto
func ReleaseTopProductGroupDto(v *TopProductGroupDto) {
	v.GroupId = ""
	v.GroupName = ""
	v.Leaf = false
	poolTopProductGroupDto.Put(v)
}
