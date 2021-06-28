package servicecenter

// PerformanceInfoWrapper 
type PerformanceInfoWrapper struct {

    // 是否有提成配置
    
    HasBonusConfig   bool `json:"has_bonus_config,omitempty" xml:"has_bonus_config,omitempty"`
    

    // 绩效数据列表
    
    PerformanceInfoList   []PerformanceInfoDTO `json:"performance_info_list,omitempty" xml:"performance_info_list,omitempty"`
    

    // 是否有授权
    
    HasAuthorize   bool `json:"has_authorize,omitempty" xml:"has_authorize,omitempty"`
    

    // 统计结束时间
    
    StatisticsEndTime   string `json:"statistics_end_time,omitempty" xml:"statistics_end_time,omitempty"`
    

    // 统计开始时间
    
    StatisticsStartTime   string `json:"statistics_start_time,omitempty" xml:"statistics_start_time,omitempty"`
    

}
