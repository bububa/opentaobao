package feedflow

// CampaignQueryDto 
type CampaignQueryDto struct {
    // 计划id
    CampaignId   int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
    // 计划名称
    CampaignName   string `json:"campaign_name,omitempty" xml:"campaign_name,omitempty"`
    // 起始位置
    Offset   int64 `json:"offset,omitempty" xml:"offset,omitempty"`
    // 每页大小
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    // 状态列表
    StatusList   []string `json:"status_list,omitempty" xml:"status_list>string,omitempty"`
}
