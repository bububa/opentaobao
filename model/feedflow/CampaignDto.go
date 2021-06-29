package feedflow

// CampaignDTo 
type CampaignDTo struct {

    // 计划id
    
    CampaignId   int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
    

    // 计划名称
    
    CampaignName   string `json:"campaign_name,omitempty" xml:"campaign_name,omitempty"`
    

    // 日预算金额，单位分
    
    DayBudget   int64 `json:"day_budget,omitempty" xml:"day_budget,omitempty"`
    

    // 投放地域
    
    LaunchAreaList   []LaunchAreaDto `json:"launch_area_list,omitempty" xml:"launch_area_list,omitempty"`
    

    // 投放时间
    
    LaunchTime   *LaunchTimeDto `json:"launch_time,omitempty" xml:"launch_time,omitempty"`
    

    // PAUSE("投放暂停"),START("投放开始"), TERMINATE("投放停止"),ABNORMAL(投放异常"),WAIT("投放等待中"),DELETE("删除")
    
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    

    // 折扣时间
    
    LaunchPeriodList   []LaunchPeriodDto `json:"launch_period_list,omitempty" xml:"launch_period_list,omitempty"`
    

}
