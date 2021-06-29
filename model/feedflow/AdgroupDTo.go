package feedflow

// AdgroupDTo 
type AdgroupDTo struct {
    // 智能出价信息
    IntelligentBid   *IntelligentBidDTO `json:"intelligent_bid,omitempty" xml:"intelligent_bid,omitempty"`
    // 商品id
    ItemId   int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
    // 单元id
    AdgroupId   int64 `json:"adgroup_id,omitempty" xml:"adgroup_id,omitempty"`
    // 单元名称
    AdgroupName   string `json:"adgroup_name,omitempty" xml:"adgroup_name,omitempty"`
    // 计划ID
    CampaignId   int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
    // 状态
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
}
