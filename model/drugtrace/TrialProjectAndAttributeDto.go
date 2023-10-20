package drugtrace

import (
	"sync"
)

// TrialProjectAndAttributeDto 结构体
type TrialProjectAndAttributeDto struct {
	// 一级药物属性
	FirstAttrDtoList []FirstAttrDto `json:"first_attr_dto_list,omitempty" xml:"first_attr_dto_list>first_attr_dto,omitempty"`
	// 项目名称
	TrialProjectName string `json:"trial_project_name,omitempty" xml:"trial_project_name,omitempty"`
	// 项目编号
	TrialProjectNo string `json:"trial_project_no,omitempty" xml:"trial_project_no,omitempty"`
}

var poolTrialProjectAndAttributeDto = sync.Pool{
	New: func() any {
		return new(TrialProjectAndAttributeDto)
	},
}

// GetTrialProjectAndAttributeDto() 从对象池中获取TrialProjectAndAttributeDto
func GetTrialProjectAndAttributeDto() *TrialProjectAndAttributeDto {
	return poolTrialProjectAndAttributeDto.Get().(*TrialProjectAndAttributeDto)
}

// ReleaseTrialProjectAndAttributeDto 释放TrialProjectAndAttributeDto
func ReleaseTrialProjectAndAttributeDto(v *TrialProjectAndAttributeDto) {
	v.FirstAttrDtoList = v.FirstAttrDtoList[:0]
	v.TrialProjectName = ""
	v.TrialProjectNo = ""
	poolTrialProjectAndAttributeDto.Put(v)
}
