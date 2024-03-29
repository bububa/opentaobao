package legalsuit

import (
	"sync"
)

// StandpointDraftDto 结构体
type StandpointDraftDto struct {
	// 口径
	DefenseCaliber string `json:"defense_caliber,omitempty" xml:"defense_caliber,omitempty"`
	// 来源
	DraftSource string `json:"draft_source,omitempty" xml:"draft_source,omitempty"`
	// 示例场景/口径描述
	StandpointDesc string `json:"standpoint_desc,omitempty" xml:"standpoint_desc,omitempty"`
	// 场景名称
	SceneName string `json:"scene_name,omitempty" xml:"scene_name,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 编号
	CaseNo string `json:"case_no,omitempty" xml:"case_no,omitempty"`
	// 类型
	CaseType string `json:"case_type,omitempty" xml:"case_type,omitempty"`
	// 提交人
	SubmitPeople string `json:"submit_people,omitempty" xml:"submit_people,omitempty"`
}

var poolStandpointDraftDto = sync.Pool{
	New: func() any {
		return new(StandpointDraftDto)
	},
}

// GetStandpointDraftDto() 从对象池中获取StandpointDraftDto
func GetStandpointDraftDto() *StandpointDraftDto {
	return poolStandpointDraftDto.Get().(*StandpointDraftDto)
}

// ReleaseStandpointDraftDto 释放StandpointDraftDto
func ReleaseStandpointDraftDto(v *StandpointDraftDto) {
	v.DefenseCaliber = ""
	v.DraftSource = ""
	v.StandpointDesc = ""
	v.SceneName = ""
	v.Remark = ""
	v.CaseNo = ""
	v.CaseType = ""
	v.SubmitPeople = ""
	poolStandpointDraftDto.Put(v)
}
