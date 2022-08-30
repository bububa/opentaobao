package scs

// CampaignQueryTopDto 结构体
type CampaignQueryTopDto struct {
	// 投放状态
	StatusList []string `json:"status_list,omitempty" xml:"status_list>string,omitempty"`
	// 场景列表
	NeedSceneList []string `json:"need_scene_list,omitempty" xml:"need_scene_list>string,omitempty"`
	// 计划id列表
	CampaignIdList []int64 `json:"campaign_id_list,omitempty" xml:"campaign_id_list>int64,omitempty"`
	// 业务线编码。场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe，爆发收割adStrategyBaoFa
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 计划名称
	CampaignName string `json:"campaign_name,omitempty" xml:"campaign_name,omitempty"`
	// 开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 计划ID
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
	// 计划状态。0暂停推广，1正在推广，9结束推广
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 预算
	DayBudget *DayBudgetTopDto `json:"day_budget,omitempty" xml:"day_budget,omitempty"`
	// 页码，从0开始
	Offset int64 `json:"offset,omitempty" xml:"offset,omitempty"`
	// 系统自动生成
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 效果归因周期，支持15，30
	Effect int64 `json:"effect,omitempty" xml:"effect,omitempty"`
}
