package alihouse

import (
	"sync"
)

// CaseFieldActivityProjectsDto 结构体
type CaseFieldActivityProjectsDto struct {
	// 活动地址
	CaseFieldActivityProjectDto []CaseFieldActivityProjectDto `json:"case_field_activity_project_dto,omitempty" xml:"case_field_activity_project_dto>case_field_activity_project_dto,omitempty"`
	// 外部活动ID
	OuterActivityId string `json:"outer_activity_id,omitempty" xml:"outer_activity_id,omitempty"`
	// 状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

var poolCaseFieldActivityProjectsDto = sync.Pool{
	New: func() any {
		return new(CaseFieldActivityProjectsDto)
	},
}

// GetCaseFieldActivityProjectsDto() 从对象池中获取CaseFieldActivityProjectsDto
func GetCaseFieldActivityProjectsDto() *CaseFieldActivityProjectsDto {
	return poolCaseFieldActivityProjectsDto.Get().(*CaseFieldActivityProjectsDto)
}

// ReleaseCaseFieldActivityProjectsDto 释放CaseFieldActivityProjectsDto
func ReleaseCaseFieldActivityProjectsDto(v *CaseFieldActivityProjectsDto) {
	v.CaseFieldActivityProjectDto = v.CaseFieldActivityProjectDto[:0]
	v.OuterActivityId = ""
	v.Status = 0
	poolCaseFieldActivityProjectsDto.Put(v)
}
