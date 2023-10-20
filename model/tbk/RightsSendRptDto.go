package tbk

import (
	"sync"
)

// RightsSendRptDto 结构体
type RightsSendRptDto struct {
	// 渠道关系id的发放数据
	RelationRptList []RightsSendRelationRptDto `json:"relation_rpt_list,omitempty" xml:"relation_rpt_list>rights_send_relation_rpt_dto,omitempty"`
	// pid的发放数据
	PidRptList []RightsSendRelationRptDto `json:"pid_rpt_list,omitempty" xml:"pid_rpt_list>rights_send_relation_rpt_dto,omitempty"`
}

var poolRightsSendRptDto = sync.Pool{
	New: func() any {
		return new(RightsSendRptDto)
	},
}

// GetRightsSendRptDto() 从对象池中获取RightsSendRptDto
func GetRightsSendRptDto() *RightsSendRptDto {
	return poolRightsSendRptDto.Get().(*RightsSendRptDto)
}

// ReleaseRightsSendRptDto 释放RightsSendRptDto
func ReleaseRightsSendRptDto(v *RightsSendRptDto) {
	v.RelationRptList = v.RelationRptList[:0]
	v.PidRptList = v.PidRptList[:0]
	poolRightsSendRptDto.Put(v)
}
