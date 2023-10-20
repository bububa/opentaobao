package legalsuit

import (
	"sync"
)

// StandpointDraftOutPutDto 结构体
type StandpointDraftOutPutDto struct {
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 创建人
	Creator string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 口径表述
	StandpointDesc string `json:"standpoint_desc,omitempty" xml:"standpoint_desc,omitempty"`
	// 建议口径
	DefenseCaliber string `json:"defense_caliber,omitempty" xml:"defense_caliber,omitempty"`
	// id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolStandpointDraftOutPutDto = sync.Pool{
	New: func() any {
		return new(StandpointDraftOutPutDto)
	},
}

// GetStandpointDraftOutPutDto() 从对象池中获取StandpointDraftOutPutDto
func GetStandpointDraftOutPutDto() *StandpointDraftOutPutDto {
	return poolStandpointDraftOutPutDto.Get().(*StandpointDraftOutPutDto)
}

// ReleaseStandpointDraftOutPutDto 释放StandpointDraftOutPutDto
func ReleaseStandpointDraftOutPutDto(v *StandpointDraftOutPutDto) {
	v.GmtCreate = ""
	v.Creator = ""
	v.StandpointDesc = ""
	v.DefenseCaliber = ""
	v.Id = 0
	poolStandpointDraftOutPutDto.Put(v)
}
