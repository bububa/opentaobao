package simba

// CiaUpdateDTO 
type CiaUpdateDTO struct {
    // 计划Id
    CampaignId   int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
    // 推广组Id
    AdgroupId   int64 `json:"adgroup_id,omitempty" xml:"adgroup_id,omitempty"`
    // 是否开启智能出价
    IsSmartBidding   bool `json:"is_smart_bidding,omitempty" xml:"is_smart_bidding,omitempty"`
    // 最高溢价比例（30-300）
    MaxPremium   int64 `json:"max_premium,omitempty" xml:"max_premium,omitempty"`
    // 出价目标（3:促进收藏加购,4:促进点击,5:促进成交）
    BidTargetType   int64 `json:"bid_target_type,omitempty" xml:"bid_target_type,omitempty"`
    // 是否自动流转（0:否，1:是）
    IsCirculation   int64 `json:"is_circulation,omitempty" xml:"is_circulation,omitempty"`
}
