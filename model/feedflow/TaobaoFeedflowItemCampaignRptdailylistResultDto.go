package feedflow

// TaobaoFeedflowItemCampaignRptdailylistResultDto 结构体
type TaobaoFeedflowItemCampaignRptdailylistResultDto struct {
	// 报表信息
	RptList []RptResultDto `json:"rpt_list,omitempty" xml:"rpt_list>rpt_result_dto,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
