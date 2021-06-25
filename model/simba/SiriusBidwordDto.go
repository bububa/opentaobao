package simba

// SiriusBidwordDto 
type SiriusBidwordDto struct {

    // 关键词无线渠道出价
    MobileBidPrice   int64 `json:"mobile_bid_price,omitempty"`

    // 关键词id
    Id   int64 `json:"id,omitempty"`

    // 推广单元id
    AdgroupId   int64 `json:"adgroup_id,omitempty"`

    // 关键词
    Word   string `json:"word,omitempty"`

    // 垃圾词状态1 为垃圾词，0 为正常词
    GarbageStatus   int64 `json:"garbage_status,omitempty"`

    // 匹配方式1 精准匹配，4广泛匹配
    MatchScope   int64 `json:"match_scope,omitempty"`

    // 推广计划id
    CampaignId   int64 `json:"campaign_id,omitempty"`

    // 关键词审核拒绝原因，当且仅当audit_status为审核拒绝时有效
    AuditReason   string `json:"audit_reason,omitempty"`

    // 审核原因0 -1 待处理，1审核通过，其他审核拒绝
    AuditStatus   int64 `json:"audit_status,omitempty"`

    // 关键词计划删除时间，仅当garbage_status=1时有效
    PlanDeleteTime   string `json:"plan_delete_time,omitempty"`

    // 关键词id
    BidwordId   int64 `json:"bidword_id,omitempty"`

}
