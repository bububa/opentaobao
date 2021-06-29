package feedflow

// TaobaoFeedflowItemCampaignRptdailylistResultDto 
type TaobaoFeedflowItemCampaignRptdailylistResultDto struct {

    // message
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    

    // 报表信息
    
    RptList   []RptResultDto `json:"rpt_list,omitempty" xml:"rpt_list,omitempty"`
    

    // 是否调用成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

}
