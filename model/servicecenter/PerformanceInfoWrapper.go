package servicecenter

// PerformanceInfoWrapper 结构体
type PerformanceInfoWrapper struct {
	// 绩效数据列表
	PerformanceInfoList []PerformanceInfoDto `json:"performance_info_list,omitempty" xml:"performance_info_list>performance_info_dto,omitempty"`
	// 统计结束时间
	StatisticsEndTime string `json:"statistics_end_time,omitempty" xml:"statistics_end_time,omitempty"`
	// 统计开始时间
	StatisticsStartTime string `json:"statistics_start_time,omitempty" xml:"statistics_start_time,omitempty"`
	// 是否有提成配置
	HasBonusConfig bool `json:"has_bonus_config,omitempty" xml:"has_bonus_config,omitempty"`
	// 是否有授权
	HasAuthorize bool `json:"has_authorize,omitempty" xml:"has_authorize,omitempty"`
}
