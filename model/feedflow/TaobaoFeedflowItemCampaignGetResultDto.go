package feedflow

// TaobaoFeedflowItemCampaignGetResultDTO 
type TaobaoFeedflowItemCampaignGetResultDTO struct {
    // 信息
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 计划信息
    Result   *CampaignDTO `json:"result,omitempty" xml:"result,omitempty"`
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
