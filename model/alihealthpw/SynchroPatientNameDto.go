package alihealthpw

import (
	"sync"
)

// SynchroPatientNameDto 结构体
type SynchroPatientNameDto struct {
	// 患者姓名
	PatientName string `json:"patient_name,omitempty" xml:"patient_name,omitempty"`
	// 修改时间
	ModifyTime string `json:"modify_time,omitempty" xml:"modify_time,omitempty"`
	// 唯一编码
	UserUniqueCode string `json:"user_unique_code,omitempty" xml:"user_unique_code,omitempty"`
	// 申请项目
	ProjectThirdId string `json:"project_third_id,omitempty" xml:"project_third_id,omitempty"`
}

var poolSynchroPatientNameDto = sync.Pool{
	New: func() any {
		return new(SynchroPatientNameDto)
	},
}

// GetSynchroPatientNameDto() 从对象池中获取SynchroPatientNameDto
func GetSynchroPatientNameDto() *SynchroPatientNameDto {
	return poolSynchroPatientNameDto.Get().(*SynchroPatientNameDto)
}

// ReleaseSynchroPatientNameDto 释放SynchroPatientNameDto
func ReleaseSynchroPatientNameDto(v *SynchroPatientNameDto) {
	v.PatientName = ""
	v.ModifyTime = ""
	v.UserUniqueCode = ""
	v.ProjectThirdId = ""
	poolSynchroPatientNameDto.Put(v)
}
