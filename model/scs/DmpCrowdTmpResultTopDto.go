package scs

import (
	"sync"
)

// DmpCrowdTmpResultTopDto 结构体
type DmpCrowdTmpResultTopDto struct {
	// group_ids
	GroupIdList []int64 `json:"group_id_list,omitempty" xml:"group_id_list>int64,omitempty"`
	// highlight
	Highlight string `json:"highlight,omitempty" xml:"highlight,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// desc
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 过期时间
	ValidDate string `json:"valid_date,omitempty" xml:"valid_date,omitempty"`
	// id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolDmpCrowdTmpResultTopDto = sync.Pool{
	New: func() any {
		return new(DmpCrowdTmpResultTopDto)
	},
}

// GetDmpCrowdTmpResultTopDto() 从对象池中获取DmpCrowdTmpResultTopDto
func GetDmpCrowdTmpResultTopDto() *DmpCrowdTmpResultTopDto {
	return poolDmpCrowdTmpResultTopDto.Get().(*DmpCrowdTmpResultTopDto)
}

// ReleaseDmpCrowdTmpResultTopDto 释放DmpCrowdTmpResultTopDto
func ReleaseDmpCrowdTmpResultTopDto(v *DmpCrowdTmpResultTopDto) {
	v.GroupIdList = v.GroupIdList[:0]
	v.Highlight = ""
	v.Name = ""
	v.Desc = ""
	v.ValidDate = ""
	v.Id = 0
	poolDmpCrowdTmpResultTopDto.Put(v)
}
