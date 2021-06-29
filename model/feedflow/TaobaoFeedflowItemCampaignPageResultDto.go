package feedflow

// TaobaoFeedflowItemCampaignPageResultDTO 
type TaobaoFeedflowItemCampaignPageResultDTO struct {
    // 信息
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 计划列表
    Results   []CampaignDTo `json:"results,omitempty" xml:"results>campaign_d_to,omitempty"`
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 符合条件的计划数量
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}
