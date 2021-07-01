package promotion

// LotteryAwardExtendDto 结构体
type LotteryAwardExtendDto struct {
	// 奖品id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 奖品类型
	AwardType int64 `json:"award_type,omitempty" xml:"award_type,omitempty"`
	// 奖品实例id
	AwardInstId string `json:"award_inst_id,omitempty" xml:"award_inst_id,omitempty"`
	// 总库存
	TotalResCount int64 `json:"total_res_count,omitempty" xml:"total_res_count,omitempty"`
	// 可用库存
	CanUseResCount int64 `json:"can_use_res_count,omitempty" xml:"can_use_res_count,omitempty"`
	// 所属活动id
	LotteryActivityId int64 `json:"lottery_activity_id,omitempty" xml:"lottery_activity_id,omitempty"`
	// 奖品显示名称
	DisplayName string `json:"display_name,omitempty" xml:"display_name,omitempty"`
	// 奖品卖家设置名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 使用门槛，分，可空
	StartFee int64 `json:"start_fee,omitempty" xml:"start_fee,omitempty"`
	// 奖品价值，分
	AwardPrice int64 `json:"award_price,omitempty" xml:"award_price,omitempty"`
	// 币种
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 币种符号
	CurrencySign string `json:"currency_sign,omitempty" xml:"currency_sign,omitempty"`
	// 奖品发放开始时间
	StartDate string `json:"start_date,omitempty" xml:"start_date,omitempty"`
	// 奖品发放结束时间
	EndDate string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// 奖品使用开始时间，可空
	UseStartDate string `json:"use_start_date,omitempty" xml:"use_start_date,omitempty"`
	// 奖品使用结束时间，可空
	UseEndDate string `json:"use_end_date,omitempty" xml:"use_end_date,omitempty"`
	// 奖品类型名称
	AwardTypeName string `json:"award_type_name,omitempty" xml:"award_type_name,omitempty"`
	// 奖品logo
	PictUrl string `json:"pict_url,omitempty" xml:"pict_url,omitempty"`
	// 使用门槛文案，可空
	Condition string `json:"condition,omitempty" xml:"condition,omitempty"`
	// 奖品详情url
	AwardDetailUrl string `json:"award_detail_url,omitempty" xml:"award_detail_url,omitempty"`
	// 奖品关联的规则
	RuleList []ExpressionRuleDto `json:"rule_list,omitempty" xml:"rule_list>expression_rule_dto,omitempty"`
}
