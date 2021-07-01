package scbp

// AccountReportOperationDto 结构体
type AccountReportOperationDto struct {
	// 计划id集合
	CampaignIds []int64 `json:"campaign_ids,omitempty" xml:"campaign_ids>int64,omitempty"`
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
