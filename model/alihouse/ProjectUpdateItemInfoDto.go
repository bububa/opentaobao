package alihouse

import (
	"sync"
)

// ProjectUpdateItemInfoDto 结构体
type ProjectUpdateItemInfoDto struct {
	// 楼盘的外部id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 楼盘描述介绍
	ProjectDescIntroduce string `json:"project_desc_introduce,omitempty" xml:"project_desc_introduce,omitempty"`
	// 外部门店ID
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
}

var poolProjectUpdateItemInfoDto = sync.Pool{
	New: func() any {
		return new(ProjectUpdateItemInfoDto)
	},
}

// GetProjectUpdateItemInfoDto() 从对象池中获取ProjectUpdateItemInfoDto
func GetProjectUpdateItemInfoDto() *ProjectUpdateItemInfoDto {
	return poolProjectUpdateItemInfoDto.Get().(*ProjectUpdateItemInfoDto)
}

// ReleaseProjectUpdateItemInfoDto 释放ProjectUpdateItemInfoDto
func ReleaseProjectUpdateItemInfoDto(v *ProjectUpdateItemInfoDto) {
	v.OuterId = ""
	v.ProjectDescIntroduce = ""
	v.OuterStoreId = ""
	poolProjectUpdateItemInfoDto.Put(v)
}
