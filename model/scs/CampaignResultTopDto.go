package scs

import (
	"sync"
)

// CampaignResultTopDto 结构体
type CampaignResultTopDto struct {
	// 单元id列表
	AdgroupIdList []int64 `json:"adgroup_id_list,omitempty" xml:"adgroup_id_list>int64,omitempty"`
	// 创意id列表
	CreativeIdList []int64 `json:"creative_id_list,omitempty" xml:"creative_id_list>int64,omitempty"`
	// 加热的商品id
	HeatItemIds []int64 `json:"heat_item_ids,omitempty" xml:"heat_item_ids>int64,omitempty"`
	// 推广计划名称
	CampaignName string `json:"campaign_name,omitempty" xml:"campaign_name,omitempty"`
	// 最高点击单价
	CpcLimit string `json:"cpc_limit,omitempty" xml:"cpc_limit,omitempty"`
	// 日预算 元
	DayBudget string `json:"day_budget,omitempty" xml:"day_budget,omitempty"`
	// 周期预算 元
	CyclicalBudget string `json:"cyclical_budget,omitempty" xml:"cyclical_budget,omitempty"`
	// 周期预算生效开始时间
	DmcBeginTime string `json:"dmc_begin_time,omitempty" xml:"dmc_begin_time,omitempty"`
	// 周期预算生效结束时间
	DmcEndTime string `json:"dmc_end_time,omitempty" xml:"dmc_end_time,omitempty"`
	// marketSceneType
	MarketSceneType string `json:"market_scene_type,omitempty" xml:"market_scene_type,omitempty"`
	// marketSceneName
	MaketSceneName string `json:"maket_scene_name,omitempty" xml:"maket_scene_name,omitempty"`
	// 活动
	ActivityId string `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 置顶时间
	TopTime string `json:"top_time,omitempty" xml:"top_time,omitempty"`
	// 结算消耗
	SettleCharge string `json:"settle_charge,omitempty" xml:"settle_charge,omitempty"`
	// 来源渠道
	SourceChannel string `json:"source_channel,omitempty" xml:"source_channel,omitempty"`
	// 出价约束的值
	ConstraintValue string `json:"constraint_value,omitempty" xml:"constraint_value,omitempty"`
	// 策略中心人群信息
	StrategyCrowdInfo string `json:"strategy_crowd_info,omitempty" xml:"strategy_crowd_info,omitempty"`
	// 计划ID
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
	// 计划组Id
	CampaignGroupId int64 `json:"campaign_group_id,omitempty" xml:"campaign_group_id,omitempty"`
	// 类型(1:CPT,2:竞价CPM,8:CPC,16:单品)
	CampaignType int64 `json:"campaign_type,omitempty" xml:"campaign_type,omitempty"`
	// 周期预算的周期天数 [3,30]
	DmcPeriod int64 `json:"dmc_period,omitempty" xml:"dmc_period,omitempty"`
	// 系统出价类型
	SystemBid int64 `json:"system_bid,omitempty" xml:"system_bid,omitempty"`
	// 是否智能出价
	AutoBid int64 `json:"auto_bid,omitempty" xml:"auto_bid,omitempty"`
	// 投放方式(1:尽快，2:平滑)
	SpeedType int64 `json:"speed_type,omitempty" xml:"speed_type,omitempty"`
	// 周期投放方式(0:尽快，1:平滑)
	PeriodSpeedType int64 `json:"period_speed_type,omitempty" xml:"period_speed_type,omitempty"`
	// 方案类型(1：智能, 2:自定义）
	SolutionType int64 `json:"solution_type,omitempty" xml:"solution_type,omitempty"`
	// marketScene
	MarketScene int64 `json:"market_scene,omitempty" xml:"market_scene,omitempty"`
	// 投放诉求（1:拓展新客, 2:打造爆款, 3: 日常销售）
	MarketAim int64 `json:"market_aim,omitempty" xml:"market_aim,omitempty"`
	// 投放状态(2:暂停,3:等待投放,5:投放中,7:投放结束, -99:异常)
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 生命周期(用来标示该计划组是否创建完成.99表示创建完成)
	LifeCycle int64 `json:"life_cycle,omitempty" xml:"life_cycle,omitempty"`
	// targetBuyer
	TargetBuyer int64 `json:"target_buyer,omitempty" xml:"target_buyer,omitempty"`
	// 来源实体id
	SourceEntityId int64 `json:"source_entity_id,omitempty" xml:"source_entity_id,omitempty"`
	// 预算类型
	DmcType int64 `json:"dmc_type,omitempty" xml:"dmc_type,omitempty"`
	// 单元间预算分配优化
	AutoDmc int64 `json:"auto_dmc,omitempty" xml:"auto_dmc,omitempty"`
	// 出价约束类型
	ConstraintType int64 `json:"constraint_type,omitempty" xml:"constraint_type,omitempty"`
	// 平均每日预算
	AvgDmc int64 `json:"avg_dmc,omitempty" xml:"avg_dmc,omitempty"`
	// 周期总预算(分)
	TotalBudget int64 `json:"total_budget,omitempty" xml:"total_budget,omitempty"`
	// abTestOpen
	AbTestOpen int64 `json:"ab_test_open,omitempty" xml:"ab_test_open,omitempty"`
	// 物料id
	MaterialId int64 `json:"material_id,omitempty" xml:"material_id,omitempty"`
	// 订单对应广告主的userid
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// dayBudgetUnlimit
	DayBudgetUnlimit bool `json:"day_budget_unlimit,omitempty" xml:"day_budget_unlimit,omitempty"`
	// channelAdzoneDiscountSwitch
	ChannelAdzoneDiscount bool `json:"channel_adzone_discount,omitempty" xml:"channel_adzone_discount,omitempty"`
}

var poolCampaignResultTopDto = sync.Pool{
	New: func() any {
		return new(CampaignResultTopDto)
	},
}

// GetCampaignResultTopDto() 从对象池中获取CampaignResultTopDto
func GetCampaignResultTopDto() *CampaignResultTopDto {
	return poolCampaignResultTopDto.Get().(*CampaignResultTopDto)
}

// ReleaseCampaignResultTopDto 释放CampaignResultTopDto
func ReleaseCampaignResultTopDto(v *CampaignResultTopDto) {
	v.AdgroupIdList = v.AdgroupIdList[:0]
	v.CreativeIdList = v.CreativeIdList[:0]
	v.HeatItemIds = v.HeatItemIds[:0]
	v.CampaignName = ""
	v.CpcLimit = ""
	v.DayBudget = ""
	v.CyclicalBudget = ""
	v.DmcBeginTime = ""
	v.DmcEndTime = ""
	v.MarketSceneType = ""
	v.MaketSceneName = ""
	v.ActivityId = ""
	v.TopTime = ""
	v.SettleCharge = ""
	v.SourceChannel = ""
	v.ConstraintValue = ""
	v.StrategyCrowdInfo = ""
	v.CampaignId = 0
	v.CampaignGroupId = 0
	v.CampaignType = 0
	v.DmcPeriod = 0
	v.SystemBid = 0
	v.AutoBid = 0
	v.SpeedType = 0
	v.PeriodSpeedType = 0
	v.SolutionType = 0
	v.MarketScene = 0
	v.MarketAim = 0
	v.Status = 0
	v.LifeCycle = 0
	v.TargetBuyer = 0
	v.SourceEntityId = 0
	v.DmcType = 0
	v.AutoDmc = 0
	v.ConstraintType = 0
	v.AvgDmc = 0
	v.TotalBudget = 0
	v.AbTestOpen = 0
	v.MaterialId = 0
	v.UserId = 0
	v.DayBudgetUnlimit = false
	v.ChannelAdzoneDiscount = false
	poolCampaignResultTopDto.Put(v)
}
