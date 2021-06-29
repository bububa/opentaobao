package feedflow

// CrowdQueryDTO 
type CrowdQueryDTO struct {
    // 定向类型
    TargetTypeList   []string `json:"target_type_list,omitempty" xml:"target_type_list>string,omitempty"`
    // 单元id
    AdgroupId   int64 `json:"adgroup_id,omitempty" xml:"adgroup_id,omitempty"`
    // 分页条件
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    // 人群id
    CrowdId   int64 `json:"crowd_id,omitempty" xml:"crowd_id,omitempty"`
    // 人群状态
    StatusList   []string `json:"status_list,omitempty" xml:"status_list>string,omitempty"`
    // 分页条件
    Offset   int64 `json:"offset,omitempty" xml:"offset,omitempty"`
    // 计划id
    CampaignId   int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
}
