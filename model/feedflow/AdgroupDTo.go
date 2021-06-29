package feedflow

// AdgroupDto 
type AdgroupDto struct {
    // 单元名称
    AdgroupName   string `json:"adgroup_name,omitempty" xml:"adgroup_name,omitempty"`
    // 资源类位表
    AdzoneList   []AdzoneBindDto `json:"adzone_list,omitempty" xml:"adzone_list>adzone_bind_dto,omitempty"`
    // 计划id
    CampaignId   int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
    // 定向人群
    CrowdList   []CrowdDto `json:"crowd_list,omitempty" xml:"crowd_list>crowd_dto,omitempty"`
    // 智能调价
    IntelligentBid   *IntelligentBidDto `json:"intelligent_bid,omitempty" xml:"intelligent_bid,omitempty"`
    // 商品id
    ItemId   int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
    // 单元id
    AdgroupId   int64 `json:"adgroup_id,omitempty" xml:"adgroup_id,omitempty"`
    // 状态
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
}
