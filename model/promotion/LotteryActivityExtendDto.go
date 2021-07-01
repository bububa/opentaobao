package promotion

// LotteryActivityExtendDto 结构体
type LotteryActivityExtendDto struct {
	// 活动ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 活动名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 状态（1有效）
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 活动开始时间
	StartDate string `json:"start_date,omitempty" xml:"start_date,omitempty"`
	// 活动结束时间
	EndDate string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// 抽奖CODE
	RaffleCode string `json:"raffle_code,omitempty" xml:"raffle_code,omitempty"`
	// 奖品列表
	AwardList []LotteryAwardExtendDto `json:"award_list,omitempty" xml:"award_list>lottery_award_extend_dto,omitempty"`
	// 活动关联的规则
	RuleList []ExpressionRuleDto `json:"rule_list,omitempty" xml:"rule_list>expression_rule_dto,omitempty"`
	// 抽奖方案列表
	SchemaList []LotterySchemaDto `json:"schema_list,omitempty" xml:"schema_list>lottery_schema_dto,omitempty"`
	// 外部关联列表
	RelationList []string `json:"relation_list,omitempty" xml:"relation_list>string,omitempty"`
}
