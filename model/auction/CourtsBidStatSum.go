package auction

// CourtsBidStatSum 
type CourtsBidStatSum struct {

    // 发拍次数
    PublishCount   int64 `json:"publish_count,omitempty"`

    // 成交金额
    HammerPrice   int64 `json:"hammer_price,omitempty"`

    // 成交件数
    HammerCount   int64 `json:"hammer_count,omitempty"`

    // 结束网拍次数
    EndCount   int64 `json:"end_count,omitempty"`

    // 结束标的件数（去重）
    EndCountDist   int64 `json:"end_count_dist,omitempty"`

    // 平均成交溢价率(万分位)
    AvgAddvPercent   int64 `json:"avg_addv_percent,omitempty"`

    // 发拍件数（去重）
    PublishCountDist   int64 `json:"publish_count_dist,omitempty"`

    // 开拍件数（去重）
    StartCountDist   int64 `json:"start_count_dist,omitempty"`

    // 出价次数
    BidCount   int64 `json:"bid_count,omitempty"`

    // 起拍价（成交标的）
    StartPrice   int64 `json:"start_price,omitempty"`

    // 意向用户数（交保数+订阅数）
    InterestCount   int64 `json:"interest_count,omitempty"`

    // 开拍数
    StartCount   int64 `json:"start_count,omitempty"`

    // 报名人数（含交保失败）
    ApplyCount   int64 `json:"apply_count,omitempty"`

}
