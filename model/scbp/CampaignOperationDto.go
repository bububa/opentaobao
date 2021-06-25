package scbp

// CampaignOperationDto 
type CampaignOperationDto struct {

    // 标题
    Title   string `json:"title,omitempty"`

    // 最小价格
    MinPrice   string `json:"min_price,omitempty"`

    // 最大价格
    MaxPrice   string `json:"max_price,omitempty"`

    // 计划状态
    OnlineStatus   int64 `json:"online_status,omitempty"`

    // 计划类型
    Type   int64 `json:"type,omitempty"`

    // 出价模式  * value =1 为智能出价      * value =2 为手动出价
    BidType   int64 `json:"bid_type,omitempty"`

    // 计划id
    Id   int64 `json:"id,omitempty"`

    // 预算
    Budget   string `json:"budget,omitempty"`

    // 开始时间
    StartTime   string `json:"start_time,omitempty"`

}
