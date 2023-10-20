package maitix

import (
	"sync"
)

// DisPerformStatusDto 结构体
type DisPerformStatusDto struct {
	// 场次id
	PerformId int64 `json:"perform_id,omitempty" xml:"perform_id,omitempty"`
	// 项目id
	ProjectId int64 `json:"project_id,omitempty" xml:"project_id,omitempty"`
	// 场次状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

var poolDisPerformStatusDto = sync.Pool{
	New: func() any {
		return new(DisPerformStatusDto)
	},
}

// GetDisPerformStatusDto() 从对象池中获取DisPerformStatusDto
func GetDisPerformStatusDto() *DisPerformStatusDto {
	return poolDisPerformStatusDto.Get().(*DisPerformStatusDto)
}

// ReleaseDisPerformStatusDto 释放DisPerformStatusDto
func ReleaseDisPerformStatusDto(v *DisPerformStatusDto) {
	v.PerformId = 0
	v.ProjectId = 0
	v.Status = 0
	poolDisPerformStatusDto.Put(v)
}
