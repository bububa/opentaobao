package promotion

// ActivityDto 结构体
type ActivityDto struct {
	// 权益列表
	BenefitList []BenefitDto `json:"benefit_list,omitempty" xml:"benefit_list>benefit_dto,omitempty"`
	// 活动名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 活动开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 活动结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 业务来源
	Source string `json:"source,omitempty" xml:"source,omitempty"`
	// 投放计划code
	StrategyCode string `json:"strategy_code,omitempty" xml:"strategy_code,omitempty"`
	// 活动状态，EFFECTIVE为生效，OFFLINE为下线
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 渠道code
	ChannelCode string `json:"channel_code,omitempty" xml:"channel_code,omitempty"`
	// 扩展字段
	Feature string `json:"feature,omitempty" xml:"feature,omitempty"`
}
