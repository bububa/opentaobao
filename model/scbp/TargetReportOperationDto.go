package scbp

import (
	"sync"
)

// TargetReportOperationDto 结构体
type TargetReportOperationDto struct {
	// crowd/region
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 开始时间(yyyy-MM-dd)
	DateBegin string `json:"date_begin,omitempty" xml:"date_begin,omitempty"`
	// 结束时间(yyyy-MM-dd)
	DateEnd string `json:"date_end,omitempty" xml:"date_end,omitempty"`
	// 营销目标
	CampaignType int64 `json:"campaign_type,omitempty" xml:"campaign_type,omitempty"`
	// 计划ID
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
}

var poolTargetReportOperationDto = sync.Pool{
	New: func() any {
		return new(TargetReportOperationDto)
	},
}

// GetTargetReportOperationDto() 从对象池中获取TargetReportOperationDto
func GetTargetReportOperationDto() *TargetReportOperationDto {
	return poolTargetReportOperationDto.Get().(*TargetReportOperationDto)
}

// ReleaseTargetReportOperationDto 释放TargetReportOperationDto
func ReleaseTargetReportOperationDto(v *TargetReportOperationDto) {
	v.Type = ""
	v.DateBegin = ""
	v.DateEnd = ""
	v.CampaignType = 0
	v.CampaignId = 0
	poolTargetReportOperationDto.Put(v)
}
