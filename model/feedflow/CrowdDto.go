package feedflow

// CrowdDto 
type CrowdDto struct {
    // 人群描述
    CrowdDesc   string `json:"crowd_desc,omitempty" xml:"crowd_desc,omitempty"`
    // 人群名称
    CrowdName   string `json:"crowd_name,omitempty" xml:"crowd_name,omitempty"`
    // 人群出价，单位：分
    Price   int64 `json:"price,omitempty" xml:"price,omitempty"`
    // 定向
    TargetLabel   *LabelDto `json:"target_label,omitempty" xml:"target_label,omitempty"`
    // 人群平均出价，单位：分
    AveragePrice   int64 `json:"average_price,omitempty" xml:"average_price,omitempty"`
    // 人群建议出价，单位：分
    SuggestPrice   int64 `json:"suggest_price,omitempty" xml:"suggest_price,omitempty"`
    // 人群id
    CrowdId   int64 `json:"crowd_id,omitempty" xml:"crowd_id,omitempty"`
    // 人群状态
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    // 计划id
    CampaignId   int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
    // 单元id
    AdgroupId   int64 `json:"adgroup_id,omitempty" xml:"adgroup_id,omitempty"`
}
