package simba

import (
	"sync"
)

// TopBidWordReportQueryVo 结构体
type TopBidWordReportQueryVo struct {
	// 聚合维度
	QueryDomains []string `json:"query_domains,omitempty" xml:"query_domains>string,omitempty"`
	// 查询指标必填，与bp页面一致，具体可用值，看返回字段中reportIndex字段，下划线改为驼峰形式
	QueryFieldInList []string `json:"query_field_in_list,omitempty" xml:"query_field_in_list>string,omitempty"`
	// 计划id筛选
	StrategyCampaignIdInList []int64 `json:"strategy_campaign_id_in_list,omitempty" xml:"strategy_campaign_id_in_list>int64,omitempty"`
	// 流量来源，资源包id筛选
	AdzonePkgIdInList []int64 `json:"adzone_pkg_id_in_list,omitempty" xml:"adzone_pkg_id_in_list>int64,omitempty"`
	// 单元id筛选
	StrategyAdgroupIdInList []int64 `json:"strategy_adgroup_id_in_list,omitempty" xml:"strategy_adgroup_id_in_list>int64,omitempty"`
	// 词包还是词，word表示词，wordPkg表示词包，不传则不过滤
	BidWordTypeInList []string `json:"bid_word_type_in_list,omitempty" xml:"bid_word_type_in_list>string,omitempty"`
	// 汇总类型 sum-汇总,day-分天,week-分周,month-分月
	SplitType string `json:"split_type,omitempty" xml:"split_type,omitempty"`
	// 归因口径，zhai-末次点击归因， mta-mta归因
	UnifyType string `json:"unify_type,omitempty" xml:"unify_type,omitempty"`
	// 开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 计划id和名称筛选项
	StrategyCampaignIdOrName string `json:"strategy_campaign_id_or_name,omitempty" xml:"strategy_campaign_id_or_name,omitempty"`
	// 单元id或者名称筛选
	StrategyAdgroupIdOrName string `json:"strategy_adgroup_id_or_name,omitempty" xml:"strategy_adgroup_id_or_name,omitempty"`
	// 词模糊筛选
	StrategyBidwordNameLike string `json:"strategy_bidword_name_like,omitempty" xml:"strategy_bidword_name_like,omitempty"`
	// 词包模糊筛选
	StrategyBidwordPkgNameLike string `json:"strategy_bidword_pkg_name_like,omitempty" xml:"strategy_bidword_pkg_name_like,omitempty"`
	// 归因周期，1、3、7、15、30
	EffectEqual int64 `json:"effect_equal,omitempty" xml:"effect_equal,omitempty"`
	// 偏移量
	Offset int64 `json:"offset,omitempty" xml:"offset,omitempty"`
	// 页面大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 是否分页
	ByPage bool `json:"by_page,omitempty" xml:"by_page,omitempty"`
}

var poolTopBidWordReportQueryVo = sync.Pool{
	New: func() any {
		return new(TopBidWordReportQueryVo)
	},
}

// GetTopBidWordReportQueryVo() 从对象池中获取TopBidWordReportQueryVo
func GetTopBidWordReportQueryVo() *TopBidWordReportQueryVo {
	return poolTopBidWordReportQueryVo.Get().(*TopBidWordReportQueryVo)
}

// ReleaseTopBidWordReportQueryVo 释放TopBidWordReportQueryVo
func ReleaseTopBidWordReportQueryVo(v *TopBidWordReportQueryVo) {
	v.QueryDomains = v.QueryDomains[:0]
	v.QueryFieldInList = v.QueryFieldInList[:0]
	v.StrategyCampaignIdInList = v.StrategyCampaignIdInList[:0]
	v.AdzonePkgIdInList = v.AdzonePkgIdInList[:0]
	v.StrategyAdgroupIdInList = v.StrategyAdgroupIdInList[:0]
	v.BidWordTypeInList = v.BidWordTypeInList[:0]
	v.SplitType = ""
	v.UnifyType = ""
	v.StartTime = ""
	v.EndTime = ""
	v.StrategyCampaignIdOrName = ""
	v.StrategyAdgroupIdOrName = ""
	v.StrategyBidwordNameLike = ""
	v.StrategyBidwordPkgNameLike = ""
	v.EffectEqual = 0
	v.Offset = 0
	v.PageSize = 0
	v.ByPage = false
	poolTopBidWordReportQueryVo.Put(v)
}
