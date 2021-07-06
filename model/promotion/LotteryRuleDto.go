package promotion

// LotteryRuleDto 结构体
type LotteryRuleDto struct {
	// 对应的规则
	TargetRules []SingleRuleDto `json:"target_rules,omitempty" xml:"target_rules>single_rule_dto,omitempty"`
	// 主体ID
	MasterId int64 `json:"master_id,omitempty" xml:"master_id,omitempty"`
}
