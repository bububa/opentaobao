package promotion

// LotteryRuleDTO 
type LotteryRuleDTO struct {
    // 主体ID
    MasterId   int64 `json:"master_id,omitempty" xml:"master_id,omitempty"`
    // 对应的规则
    TargetRules   []SingleRuleDTO `json:"target_rules,omitempty" xml:"target_rules>single_rule_dto,omitempty"`
}
