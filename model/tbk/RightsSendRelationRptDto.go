package tbk

import (
	"sync"
)

// RightsSendRelationRptDto 结构体
type RightsSendRelationRptDto struct {
	// 日期
	BizDate string `json:"biz_date,omitempty" xml:"biz_date,omitempty"`
	// 渠道关系id关联的pid
	Pid string `json:"pid,omitempty" xml:"pid,omitempty"`
	// 渠道关系id
	RelationId int64 `json:"relation_id,omitempty" xml:"relation_id,omitempty"`
	// 红包发放数量
	FundNum int64 `json:"fund_num,omitempty" xml:"fund_num,omitempty"`
	// 红包使用次数
	UseNum int64 `json:"use_num,omitempty" xml:"use_num,omitempty"`
}

var poolRightsSendRelationRptDto = sync.Pool{
	New: func() any {
		return new(RightsSendRelationRptDto)
	},
}

// GetRightsSendRelationRptDto() 从对象池中获取RightsSendRelationRptDto
func GetRightsSendRelationRptDto() *RightsSendRelationRptDto {
	return poolRightsSendRelationRptDto.Get().(*RightsSendRelationRptDto)
}

// ReleaseRightsSendRelationRptDto 释放RightsSendRelationRptDto
func ReleaseRightsSendRelationRptDto(v *RightsSendRelationRptDto) {
	v.BizDate = ""
	v.Pid = ""
	v.RelationId = 0
	v.FundNum = 0
	v.UseNum = 0
	poolRightsSendRelationRptDto.Put(v)
}
