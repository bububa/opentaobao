package drugtrace

import (
	"sync"
)

// TrialProjectDto 结构体
type TrialProjectDto struct {
	// 药品信息
	DrugDtoList []SubType `json:"drug_dto_list,omitempty" xml:"drug_dto_list>sub_type,omitempty"`
	// 项目名称
	TrialProjectName string `json:"trial_project_name,omitempty" xml:"trial_project_name,omitempty"`
	// 项目编号
	TrialProjectNo string `json:"trial_project_no,omitempty" xml:"trial_project_no,omitempty"`
}

var poolTrialProjectDto = sync.Pool{
	New: func() any {
		return new(TrialProjectDto)
	},
}

// GetTrialProjectDto() 从对象池中获取TrialProjectDto
func GetTrialProjectDto() *TrialProjectDto {
	return poolTrialProjectDto.Get().(*TrialProjectDto)
}

// ReleaseTrialProjectDto 释放TrialProjectDto
func ReleaseTrialProjectDto(v *TrialProjectDto) {
	v.DrugDtoList = v.DrugDtoList[:0]
	v.TrialProjectName = ""
	v.TrialProjectNo = ""
	poolTrialProjectDto.Put(v)
}
