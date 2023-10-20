package scs

import (
	"sync"
)

// DmpCrowdResultTopDto 结构体
type DmpCrowdResultTopDto struct {
	// 人群名
	CrowdName string `json:"crowd_name,omitempty" xml:"crowd_name,omitempty"`
	// 过期时间
	ValidDate string `json:"valid_date,omitempty" xml:"valid_date,omitempty"`
	// 创建时间
	Createtime string `json:"createtime,omitempty" xml:"createtime,omitempty"`
	// 更新时间
	Updatetime string `json:"updatetime,omitempty" xml:"updatetime,omitempty"`
	// 人群id
	CrowdId int64 `json:"crowd_id,omitempty" xml:"crowd_id,omitempty"`
	// test
	Lookalike int64 `json:"lookalike,omitempty" xml:"lookalike,omitempty"`
	// 人数预估
	Coverage int64 `json:"coverage,omitempty" xml:"coverage,omitempty"`
	// 状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 状态
	FullStatus int64 `json:"full_status,omitempty" xml:"full_status,omitempty"`
	// 单元ID
	GroupId int64 `json:"group_id,omitempty" xml:"group_id,omitempty"`
	// storage
	StorageType int64 `json:"storage_type,omitempty" xml:"storage_type,omitempty"`
}

var poolDmpCrowdResultTopDto = sync.Pool{
	New: func() any {
		return new(DmpCrowdResultTopDto)
	},
}

// GetDmpCrowdResultTopDto() 从对象池中获取DmpCrowdResultTopDto
func GetDmpCrowdResultTopDto() *DmpCrowdResultTopDto {
	return poolDmpCrowdResultTopDto.Get().(*DmpCrowdResultTopDto)
}

// ReleaseDmpCrowdResultTopDto 释放DmpCrowdResultTopDto
func ReleaseDmpCrowdResultTopDto(v *DmpCrowdResultTopDto) {
	v.CrowdName = ""
	v.ValidDate = ""
	v.Createtime = ""
	v.Updatetime = ""
	v.CrowdId = 0
	v.Lookalike = 0
	v.Coverage = 0
	v.Status = 0
	v.FullStatus = 0
	v.GroupId = 0
	v.StorageType = 0
	poolDmpCrowdResultTopDto.Put(v)
}
