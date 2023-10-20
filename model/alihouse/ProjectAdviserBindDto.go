package alihouse

import (
	"sync"
)

// ProjectAdviserBindDto 结构体
type ProjectAdviserBindDto struct {
	// 外部顾问id
	OuterConsultantId string `json:"outer_consultant_id,omitempty" xml:"outer_consultant_id,omitempty"`
	// 外部楼盘id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 外部门店ID
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
}

var poolProjectAdviserBindDto = sync.Pool{
	New: func() any {
		return new(ProjectAdviserBindDto)
	},
}

// GetProjectAdviserBindDto() 从对象池中获取ProjectAdviserBindDto
func GetProjectAdviserBindDto() *ProjectAdviserBindDto {
	return poolProjectAdviserBindDto.Get().(*ProjectAdviserBindDto)
}

// ReleaseProjectAdviserBindDto 释放ProjectAdviserBindDto
func ReleaseProjectAdviserBindDto(v *ProjectAdviserBindDto) {
	v.OuterConsultantId = ""
	v.OuterId = ""
	v.OuterStoreId = ""
	poolProjectAdviserBindDto.Put(v)
}
