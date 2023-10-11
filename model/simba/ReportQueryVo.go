package simba

// ReportQueryVo 结构体
type ReportQueryVo struct {
	// 场景筛选code
	BizCodeInList []string `json:"biz_code_in_list,omitempty" xml:"biz_code_in_list>string,omitempty"`
	// 优化目标筛选
	StrategyOptimizeTargetInList []int64 `json:"strategy_optimize_target_in_list,omitempty" xml:"strategy_optimize_target_in_list>int64,omitempty"`
	// 计划id筛选
	StrategyCampaignIdInList []int64 `json:"strategy_campaign_id_in_list,omitempty" xml:"strategy_campaign_id_in_list>int64,omitempty"`
	// 单元id筛选
	StrategyAdgroupIdInList []int64 `json:"strategy_adgroup_id_in_list,omitempty" xml:"strategy_adgroup_id_in_list>int64,omitempty"`
	// 主体类型筛选，ITEM_PRIVATE_MINIL-橱窗，SHOP-店铺，USER_DEFINE_URL-自定义
	SubPromotionTypes []string `json:"sub_promotion_types,omitempty" xml:"sub_promotion_types>string,omitempty"`
	// 宝贝id筛选
	ItemIds []int64 `json:"item_ids,omitempty" xml:"item_ids>int64,omitempty"`
	// 流量来源资源包id筛选
	AdzonePkgIdInList []int64 `json:"adzone_pkg_id_in_list,omitempty" xml:"adzone_pkg_id_in_list>int64,omitempty"`
	// 词包还是词
	BidWordTypeInList []string `json:"bid_word_type_in_list,omitempty" xml:"bid_word_type_in_list>string,omitempty"`
	// 省份筛选
	ProvinceIdInList []int64 `json:"province_id_in_list,omitempty" xml:"province_id_in_list>int64,omitempty"`
	// 聚合维度，可用值和选择对应报表类型一致
	QueryDomains []string `json:"query_domains,omitempty" xml:"query_domains>string,omitempty"`
	// 需要下载的指标数据
	QueryFieldInList []string `json:"query_field_in_list,omitempty" xml:"query_field_in_list>string,omitempty"`
	// 汇总类型 sum-汇总,hour-分时,day-分天,week-分周,month-分月
	SplitType string `json:"split_type,omitempty" xml:"split_type,omitempty"`
	// 归因口径，zhai-末次点击归因， mta-mta归因
	UnifyType string `json:"unify_type,omitempty" xml:"unify_type,omitempty"`
	// 开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 单元id或者名称筛选
	StrategyAdgroupIdOrName string `json:"strategy_adgroup_id_or_name,omitempty" xml:"strategy_adgroup_id_or_name,omitempty"`
	// 人群名称模糊过滤
	StrategyTargetTitleLike string `json:"strategy_target_title_like,omitempty" xml:"strategy_target_title_like,omitempty"`
	// 仅创意报表有效,创意筛选
	StrategyCreativeIdOrName string `json:"strategy_creative_id_or_name,omitempty" xml:"strategy_creative_id_or_name,omitempty"`
	// 词模糊筛选
	StrategyBidwordNameLike string `json:"strategy_bidword_name_like,omitempty" xml:"strategy_bidword_name_like,omitempty"`
	// 词包模糊筛选
	StrategyBidwordPkgNameLike string `json:"strategy_bidword_pkg_name_like,omitempty" xml:"strategy_bidword_pkg_name_like,omitempty"`
	// 报表异步下载文件名称，下载时必填
	ExcelName string `json:"excel_name,omitempty" xml:"excel_name,omitempty"`
	// 下载指标设置， all-全部指标，selected-下载传入指标
	FieldType string `json:"field_type,omitempty" xml:"field_type,omitempty"`
	// 报表下载需要的参数，必填, report_frame_account-账户,report_frame_campaign-计划,report_frame_adgroup-单元,report_frame_bidword-关键词,report_frame_crowd-人群,report_frame_item_promotion-宝贝主体,report_frame_other_promotion-其他主体,report_frame_creative-创意,report_frame_area-地域
	ParentAdcName string `json:"parent_adc_name,omitempty" xml:"parent_adc_name,omitempty"`
	// 归因周期，1、3、7、15、30
	EffectEqual int64 `json:"effect_equal,omitempty" xml:"effect_equal,omitempty"`
	// 偏移量
	Offset int64 `json:"offset,omitempty" xml:"offset,omitempty"`
	// 页面大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}
