package feedflow

// TaobaoFeedflowItemCampaignDaybudgetResultDto 
type TaobaoFeedflowItemCampaignDaybudgetResultDto struct {
    // 信息
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 预算总额
    Result   int64 `json:"result,omitempty" xml:"result,omitempty"`
    // 成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
