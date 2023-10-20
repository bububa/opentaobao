package scs

import (
	"sync"
)

// CampaignSolutionTopDto 结构体
type CampaignSolutionTopDto struct {
	// 新达摩盘精选人群信息
	NewDmpTemplateCrowd []NewDmpTemplateCrowdTopDto `json:"new_dmp_template_crowd,omitempty" xml:"new_dmp_template_crowd>new_dmp_template_crowd_top_dto,omitempty"`
	// 计划名称
	CampaignName string `json:"campaign_name,omitempty" xml:"campaign_name,omitempty"`
	// 推广计划的计费类型，套餐包为order，持续推广为daily
	PromotionModel string `json:"promotion_model,omitempty" xml:"promotion_model,omitempty"`
	// 持续推广计划需要填的预算
	DayBudget *DayBudgetTopDto `json:"day_budget,omitempty" xml:"day_budget,omitempty"`
	// 生命周期
	LifeCycle int64 `json:"life_cycle,omitempty" xml:"life_cycle,omitempty"`
	// 场景相关信息
	Marketing *MarketingTopDto `json:"marketing,omitempty" xml:"marketing,omitempty"`
	// 计划时间
	LaunchTime *LaunchTimeTopDto `json:"launch_time,omitempty" xml:"launch_time,omitempty"`
	// 计划策略信息
	AdStrategyInfo *AdStrategyInfoTopDto `json:"ad_strategy_info,omitempty" xml:"ad_strategy_info,omitempty"`
}

var poolCampaignSolutionTopDto = sync.Pool{
	New: func() any {
		return new(CampaignSolutionTopDto)
	},
}

// GetCampaignSolutionTopDto() 从对象池中获取CampaignSolutionTopDto
func GetCampaignSolutionTopDto() *CampaignSolutionTopDto {
	return poolCampaignSolutionTopDto.Get().(*CampaignSolutionTopDto)
}

// ReleaseCampaignSolutionTopDto 释放CampaignSolutionTopDto
func ReleaseCampaignSolutionTopDto(v *CampaignSolutionTopDto) {
	v.NewDmpTemplateCrowd = v.NewDmpTemplateCrowd[:0]
	v.CampaignName = ""
	v.PromotionModel = ""
	v.DayBudget = nil
	v.LifeCycle = 0
	v.Marketing = nil
	v.LaunchTime = nil
	v.AdStrategyInfo = nil
	poolCampaignSolutionTopDto.Put(v)
}
