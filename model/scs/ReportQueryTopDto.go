package scs

import (
	"sync"
)

// ReportQueryTopDto 结构体
type ReportQueryTopDto struct {
	// 查询日期
	LogDateList []string `json:"log_date_list,omitempty" xml:"log_date_list>string,omitempty"`
	// 计划列表
	CampaignIdList []int64 `json:"campaign_id_list,omitempty" xml:"campaign_id_list>int64,omitempty"`
	// 产品线列表
	LaunchProductIdList []int64 `json:"launch_product_id_list,omitempty" xml:"launch_product_id_list>int64,omitempty"`
	// 人群列表
	WhiteCrowdIdList []int64 `json:"white_crowd_id_list,omitempty" xml:"white_crowd_id_list>int64,omitempty"`
	// 开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 效果转化类型
	EffectType string `json:"effect_type,omitempty" xml:"effect_type,omitempty"`
	// 宽窄口径（kuan/zhai）
	UnifyType string `json:"unify_type,omitempty" xml:"unify_type,omitempty"`
	// 效果转化天数
	Effect int64 `json:"effect,omitempty" xml:"effect,omitempty"`
	// 人群id
	CrowdId int64 `json:"crowd_id,omitempty" xml:"crowd_id,omitempty"`
	// 是否查询udf
	QueryUdf bool `json:"query_udf,omitempty" xml:"query_udf,omitempty"`
	// 是否场景
	StrategyScene bool `json:"strategy_scene,omitempty" xml:"strategy_scene,omitempty"`
}

var poolReportQueryTopDto = sync.Pool{
	New: func() any {
		return new(ReportQueryTopDto)
	},
}

// GetReportQueryTopDto() 从对象池中获取ReportQueryTopDto
func GetReportQueryTopDto() *ReportQueryTopDto {
	return poolReportQueryTopDto.Get().(*ReportQueryTopDto)
}

// ReleaseReportQueryTopDto 释放ReportQueryTopDto
func ReleaseReportQueryTopDto(v *ReportQueryTopDto) {
	v.LogDateList = v.LogDateList[:0]
	v.CampaignIdList = v.CampaignIdList[:0]
	v.LaunchProductIdList = v.LaunchProductIdList[:0]
	v.WhiteCrowdIdList = v.WhiteCrowdIdList[:0]
	v.StartTime = ""
	v.EndTime = ""
	v.EffectType = ""
	v.UnifyType = ""
	v.Effect = 0
	v.CrowdId = 0
	v.QueryUdf = false
	v.StrategyScene = false
	poolReportQueryTopDto.Put(v)
}
