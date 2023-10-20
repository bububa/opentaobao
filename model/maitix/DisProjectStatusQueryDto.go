package maitix

import (
	"sync"
)

// DisProjectStatusQueryDto 结构体
type DisProjectStatusQueryDto struct {
	// 项目id
	ProjectId int64 `json:"project_id,omitempty" xml:"project_id,omitempty"`
	// 是否查询对应场次的状态
	QueryPerformStatus bool `json:"query_perform_status,omitempty" xml:"query_perform_status,omitempty"`
}

var poolDisProjectStatusQueryDto = sync.Pool{
	New: func() any {
		return new(DisProjectStatusQueryDto)
	},
}

// GetDisProjectStatusQueryDto() 从对象池中获取DisProjectStatusQueryDto
func GetDisProjectStatusQueryDto() *DisProjectStatusQueryDto {
	return poolDisProjectStatusQueryDto.Get().(*DisProjectStatusQueryDto)
}

// ReleaseDisProjectStatusQueryDto 释放DisProjectStatusQueryDto
func ReleaseDisProjectStatusQueryDto(v *DisProjectStatusQueryDto) {
	v.ProjectId = 0
	v.QueryPerformStatus = false
	poolDisProjectStatusQueryDto.Put(v)
}
