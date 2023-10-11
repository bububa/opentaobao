package simba

// TopAccountReportQueryVo 结构体
type TopAccountReportQueryVo struct {
	// 聚合维度可以传空的list，date-时间，scene-营销场景
	QueryDomains []string `json:"query_domains,omitempty" xml:"query_domains>string,omitempty"`
	// 查询指标必填，与bp页面一致，具体可用值，看返回字段中reportIndex字段, 下划线改为驼峰形式
	QueryFieldInList []string `json:"query_field_in_list,omitempty" xml:"query_field_in_list>string,omitempty"`
	// 营销场景code，具体code见 account.get.can.use.bizcode 此api返回
	BizCodeInList []string `json:"biz_code_in_list,omitempty" xml:"biz_code_in_list>string,omitempty"`
	// 汇总类型 sum-汇总,hour-分时,day-分天,week-分周,month-分月
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
	// 是否分页
	ByPage bool `json:"by_page,omitempty" xml:"by_page,omitempty"`
}
