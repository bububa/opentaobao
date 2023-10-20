package scbp

import (
	"sync"
)

// AccountReportOperationDto 结构体
type AccountReportOperationDto struct {
	// 开始时间(yyyy-MM-dd)
	DateBegin string `json:"date_begin,omitempty" xml:"date_begin,omitempty"`
	// 结束时间(yyyy-MM-dd)
	DateEnd string `json:"date_end,omitempty" xml:"date_end,omitempty"`
	// 时间段
	DateRange int64 `json:"date_range,omitempty" xml:"date_range,omitempty"`
	// campaignType
	CampaignType int64 `json:"campaign_type,omitempty" xml:"campaign_type,omitempty"`
	// campaignId
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
}

var poolAccountReportOperationDto = sync.Pool{
	New: func() any {
		return new(AccountReportOperationDto)
	},
}

// GetAccountReportOperationDto() 从对象池中获取AccountReportOperationDto
func GetAccountReportOperationDto() *AccountReportOperationDto {
	return poolAccountReportOperationDto.Get().(*AccountReportOperationDto)
}

// ReleaseAccountReportOperationDto 释放AccountReportOperationDto
func ReleaseAccountReportOperationDto(v *AccountReportOperationDto) {
	v.DateBegin = ""
	v.DateEnd = ""
	v.DateRange = 0
	v.CampaignType = 0
	v.CampaignId = 0
	poolAccountReportOperationDto.Put(v)
}
