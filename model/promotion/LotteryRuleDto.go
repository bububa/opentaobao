package promotion

// LotteryRuleDto 
/* model for simplify = false
type LotteryRuleDto struct {

    // 主体ID
    
    MasterId   int64 `json:"master_id,omitempty"`
    

    // 对应的规则
    
    TargetRules  struct {
        SingleRuleDto  []SingleRuleDto `json:"single_rule_dto,omitempty"`
    } `json:"target_rules,omitempty"`
    

}
*/

// LotteryRuleDto 
type LotteryRuleDto struct {

    // 主体ID
    MasterId   int64 `json:"master_id,omitempty"`

    // 对应的规则
    TargetRules   []SingleRuleDto `json:"target_rules,omitempty"`

}
