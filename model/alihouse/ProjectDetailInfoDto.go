package alihouse

import (
	"sync"
)

// ProjectDetailInfoDto 结构体
type ProjectDetailInfoDto struct {
	// 外部楼盘ID
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
}

var poolProjectDetailInfoDto = sync.Pool{
	New: func() any {
		return new(ProjectDetailInfoDto)
	},
}

// GetProjectDetailInfoDto() 从对象池中获取ProjectDetailInfoDto
func GetProjectDetailInfoDto() *ProjectDetailInfoDto {
	return poolProjectDetailInfoDto.Get().(*ProjectDetailInfoDto)
}

// ReleaseProjectDetailInfoDto 释放ProjectDetailInfoDto
func ReleaseProjectDetailInfoDto(v *ProjectDetailInfoDto) {
	v.OuterId = ""
	poolProjectDetailInfoDto.Put(v)
}
