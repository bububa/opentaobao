package scbp

// TopP4pBasicQuickCampaign 
/* model for simplify = false
type TopP4pBasicQuickCampaign struct {

    // 出价区间-上限(一位小数，不低于3.0)
    
    MaxPrice   string `json:"max_price,omitempty"`
    

    // 出价区间-下限(一位小数)
    
    MinPrice   string `json:"min_price,omitempty"`
    

    // 计划每日预算
    
    Budget   int64 `json:"budget,omitempty"`
    

    // 计划标题
    
    Title   string `json:"title,omitempty"`
    

    // 推广计划id
    
    CampaignId   int64 `json:"campaign_id,omitempty"`
    

}
*/

// TopP4pBasicQuickCampaign 
type TopP4pBasicQuickCampaign struct {

    // 出价区间-上限(一位小数，不低于3.0)
    MaxPrice   string `json:"max_price,omitempty"`

    // 出价区间-下限(一位小数)
    MinPrice   string `json:"min_price,omitempty"`

    // 计划每日预算
    Budget   int64 `json:"budget,omitempty"`

    // 计划标题
    Title   string `json:"title,omitempty"`

    // 推广计划id
    CampaignId   int64 `json:"campaign_id,omitempty"`

}
