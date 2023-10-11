package promotion

// BenefitSelectorQuery 结构体
type BenefitSelectorQuery struct {
	// 权益类型（步骤2,3必填）
	BenefitType string `json:"benefit_type,omitempty" xml:"benefit_type,omitempty"`
	// 权益id（步骤3必填）
	BenefitId int64 `json:"benefit_id,omitempty" xml:"benefit_id,omitempty"`
	// 新权益类型id
	ConfigId int64 `json:"config_id,omitempty" xml:"config_id,omitempty"`
	// 需要过滤的option
	ExcludeOptions int64 `json:"exclude_options,omitempty" xml:"exclude_options,omitempty"`
	// 选择器步骤可选1,2,3，1：展示卖家能够选择的权益类型2：必填BenefitType，展示卖家该种类型下的权益3：必填BenefitType和benefitId，展示卖家该权益的详情
	Step int64 `json:"step,omitempty" xml:"step,omitempty"`
	// 选择器步骤选择2时，查询指定类型权益列表分页查询参数
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 选择器步骤选择2时，查询指定类型权益列表分页查询参数
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 传true时列表进行分页查询，分页查询时不会对发放完的模板过滤；传false时分页查询的current_page和page_size不传，接口语义和之前一致。
	PageQueryRequest bool `json:"page_query_request,omitempty" xml:"page_query_request,omitempty"`
}
