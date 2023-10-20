package alihouse

import (
	"sync"
)

// ProjectAdviserBindResultDto 结构体
type ProjectAdviserBindResultDto struct {
	// 外部顾问id
	OuterConsultantId string `json:"outer_consultant_id,omitempty" xml:"outer_consultant_id,omitempty"`
	// 外部楼盘id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 失败原因
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

var poolProjectAdviserBindResultDto = sync.Pool{
	New: func() any {
		return new(ProjectAdviserBindResultDto)
	},
}

// GetProjectAdviserBindResultDto() 从对象池中获取ProjectAdviserBindResultDto
func GetProjectAdviserBindResultDto() *ProjectAdviserBindResultDto {
	return poolProjectAdviserBindResultDto.Get().(*ProjectAdviserBindResultDto)
}

// ReleaseProjectAdviserBindResultDto 释放ProjectAdviserBindResultDto
func ReleaseProjectAdviserBindResultDto(v *ProjectAdviserBindResultDto) {
	v.OuterConsultantId = ""
	v.OuterId = ""
	v.Code = ""
	v.Message = ""
	poolProjectAdviserBindResultDto.Put(v)
}
