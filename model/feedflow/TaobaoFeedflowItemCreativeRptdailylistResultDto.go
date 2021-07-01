package feedflow

// TaobaoFeedflowItemCreativeRptdailylistResultDto 结构体
type TaobaoFeedflowItemCreativeRptdailylistResultDto struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 报表结果
	RptList []RptResultDto `json:"rpt_list,omitempty" xml:"rpt_list>rpt_result_dto,omitempty"`
	// 是否调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
