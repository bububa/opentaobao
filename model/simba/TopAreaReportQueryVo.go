package simba

import (
	"sync"
)

// TopAreaReportQueryVo 结构体
type TopAreaReportQueryVo struct {
	// 聚合维度，province-省份，date-时间
	QueryDomains []string `json:"query_domains,omitempty" xml:"query_domains>string,omitempty"`
	// 查询指标必填，与bp页面一致，具体可用值，看返回字段中reportIndex字段，下划线改为驼峰形式
	QueryFieldInList []string `json:"query_field_in_list,omitempty" xml:"query_field_in_list>string,omitempty"`
	// 营销场景code，具体code见 account.get.can.use.bizcode 此api返回
	BizCodeInList []string `json:"biz_code_in_list,omitempty" xml:"biz_code_in_list>string,omitempty"`
	// 省份筛选
	ProvinceIdInList []int64 `json:"province_id_in_list,omitempty" xml:"province_id_in_list>int64,omitempty"`
	// 优化目标筛选
	StrategyOptimizeTargetInList []int64 `json:"strategy_optimize_target_in_list,omitempty" xml:"strategy_optimize_target_in_list>int64,omitempty"`
	// 计划id筛选
	StrategyCampaignIdInList []int64 `json:"strategy_campaign_id_in_list,omitempty" xml:"strategy_campaign_id_in_list>int64,omitempty"`
	// 流量来源，资源包id筛选
	AdzonePkgIdInList []int64 `json:"adzone_pkg_id_in_list,omitempty" xml:"adzone_pkg_id_in_list>int64,omitempty"`
	// 汇总类型 sum-汇总,day-分天,week-分周,month-分月
	SplitType string `json:"split_type,omitempty" xml:"split_type,omitempty"`
	// 归因口径，zhai-末次点击归因， mta-mta归因
	UnifyType string `json:"unify_type,omitempty" xml:"unify_type,omitempty"`
	// 开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 归因周期，1、3、7、15、30
	EffectEqual int64 `json:"effect_equal,omitempty" xml:"effect_equal,omitempty"`
	// 偏移量
	Offset int64 `json:"offset,omitempty" xml:"offset,omitempty"`
	// 页面大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 是否分页，true-是，false-否
	ByPage bool `json:"by_page,omitempty" xml:"by_page,omitempty"`
}

var poolTopAreaReportQueryVo = sync.Pool{
	New: func() any {
		return new(TopAreaReportQueryVo)
	},
}

// GetTopAreaReportQueryVo() 从对象池中获取TopAreaReportQueryVo
func GetTopAreaReportQueryVo() *TopAreaReportQueryVo {
	return poolTopAreaReportQueryVo.Get().(*TopAreaReportQueryVo)
}

// ReleaseTopAreaReportQueryVo 释放TopAreaReportQueryVo
func ReleaseTopAreaReportQueryVo(v *TopAreaReportQueryVo) {
	v.QueryDomains = v.QueryDomains[:0]
	v.QueryFieldInList = v.QueryFieldInList[:0]
	v.BizCodeInList = v.BizCodeInList[:0]
	v.ProvinceIdInList = v.ProvinceIdInList[:0]
	v.StrategyOptimizeTargetInList = v.StrategyOptimizeTargetInList[:0]
	v.StrategyCampaignIdInList = v.StrategyCampaignIdInList[:0]
	v.AdzonePkgIdInList = v.AdzonePkgIdInList[:0]
	v.SplitType = ""
	v.UnifyType = ""
	v.StartTime = ""
	v.EndTime = ""
	v.EffectEqual = 0
	v.Offset = 0
	v.PageSize = 0
	v.ByPage = false
	poolTopAreaReportQueryVo.Put(v)
}
