package drugtrace

import (
	"sync"
)

// TrialProjectInfoDto 结构体
type TrialProjectInfoDto struct {
	// 项目和药物属性
	ProjectAndAttributeDtoList []TrialProjectAndAttributeDto `json:"project_and_attribute_dto_list,omitempty" xml:"project_and_attribute_dto_list>trial_project_and_attribute_dto,omitempty"`
	// 药品信息
	DrugDtoList []SubType `json:"drug_dto_list,omitempty" xml:"drug_dto_list>sub_type,omitempty"`
}

var poolTrialProjectInfoDto = sync.Pool{
	New: func() any {
		return new(TrialProjectInfoDto)
	},
}

// GetTrialProjectInfoDto() 从对象池中获取TrialProjectInfoDto
func GetTrialProjectInfoDto() *TrialProjectInfoDto {
	return poolTrialProjectInfoDto.Get().(*TrialProjectInfoDto)
}

// ReleaseTrialProjectInfoDto 释放TrialProjectInfoDto
func ReleaseTrialProjectInfoDto(v *TrialProjectInfoDto) {
	v.ProjectAndAttributeDtoList = v.ProjectAndAttributeDtoList[:0]
	v.DrugDtoList = v.DrugDtoList[:0]
	poolTrialProjectInfoDto.Put(v)
}
