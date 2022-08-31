package scbp

// KeywordReportOperationDto 结构体
type KeywordReportOperationDto struct {
	// 精确搜索关键词
	KeywordList []string `json:"keyword_list,omitempty" xml:"keyword_list>string,omitempty"`
	// 效果报告模糊搜索关键词
	Keyword string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// 获取明细数据（&#34;true&#34;/&#34;false&#34;）,如果为&#34;true&#34;则为明细数据
	GetDetailData string `json:"get_detail_data,omitempty" xml:"get_detail_data,omitempty"`
	// 开始时间(yyyy-MM-dd)
	DateBegin string `json:"date_begin,omitempty" xml:"date_begin,omitempty"`
	// 结束时间(yyyy-MM-dd)
	DateEnd string `json:"date_end,omitempty" xml:"date_end,omitempty"`
	// 排序字段(impr/click/ctr/cost/cpc)
	OrderField string `json:"order_field,omitempty" xml:"order_field,omitempty"`
	// 排序方式(asc/desc)
	OrderType string `json:"order_type,omitempty" xml:"order_type,omitempty"`
	// 时间段
	DateRange int64 `json:"date_range,omitempty" xml:"date_range,omitempty"`
	// 营销目标
	CampaignType int64 `json:"campaign_type,omitempty" xml:"campaign_type,omitempty"`
	// 计划ID
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
	// 分页数量
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 页码
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
}
