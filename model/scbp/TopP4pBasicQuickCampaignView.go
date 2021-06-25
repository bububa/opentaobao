package scbp

// TopP4pBasicQuickCampaignView 
type TopP4pBasicQuickCampaignView struct {

    // 屏蔽词列表
    ForbiddenWords   []String `json:"forbidden_words,omitempty"`

    // 计划每日预算
    Budget   int64 `json:"budget,omitempty"`

    // 出价区间-上限
    MaxPrice   string `json:"max_price,omitempty"`

    // 出价区间-下限
    MinPrice   string `json:"min_price,omitempty"`

    // 推广商品数量
    ProductCount   int64 `json:"product_count,omitempty"`

    // 计划状态(0暂停、1推广中、-1点爆)
    Status   int64 `json:"status,omitempty"`

    // 计划标题
    Title   string `json:"title,omitempty"`

    // 定向推广计划ID
    CampaignId   int64 `json:"campaign_id,omitempty"`

}
