package feedflow

// TaobaoFeedflowItemAdzoneRptdailylistResultDTO 
type TaobaoFeedflowItemAdzoneRptdailylistResultDTO struct {
    // message
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 报表信息
    RptList   []RptResultDTO `json:"rpt_list,omitempty" xml:"rpt_list>rpt_result_dto,omitempty"`
    // 是否调用成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
