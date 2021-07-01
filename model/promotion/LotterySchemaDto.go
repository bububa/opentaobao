package promotion

// LotterySchemaDto 
type LotterySchemaDto struct {
    // 方案id
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    // 修改时间
    GmtModified   string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
    // 方案名称
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 概率
    Probability   string `json:"probability,omitempty" xml:"probability,omitempty"`
    // 抽奖活动id
    LotteryActivityId   int64 `json:"lottery_activity_id,omitempty" xml:"lottery_activity_id,omitempty"`
    // 方案关联的规则列表
    RuleList   []ExpressionRuleDto `json:"rule_list,omitempty" xml:"rule_list>expression_rule_dto,omitempty"`
    // 方案关联的奖品
    AwardList   []LotteryAwardRelDto `json:"award_list,omitempty" xml:"award_list>lottery_award_rel_dto,omitempty"`
    // 业务标
    Flag   int64 `json:"flag,omitempty" xml:"flag,omitempty"`
}
