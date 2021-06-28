package promotion

// LotteryRuleCreateDto 
/* model for simplify = false
type LotteryRuleCreateDto struct {

    // 主体类型（1活动2奖品3方案）
    
    MasterType   int64 `json:"master_type,omitempty"`
    

    // 批量规则列表
    
    RuleList  struct {
        LotteryRuleDto  []LotteryRuleDto `json:"lottery_rule_dto,omitempty"`
    } `json:"rule_list,omitempty"`
    

    // 抽奖活动ID
    
    LotteryActivityId   int64 `json:"lottery_activity_id,omitempty"`
    

}
*/

// LotteryRuleCreateDto 
type LotteryRuleCreateDto struct {

    // 主体类型（1活动2奖品3方案）
    MasterType   int64 `json:"master_type,omitempty"`

    // 批量规则列表
    RuleList   []LotteryRuleDto `json:"rule_list,omitempty"`

    // 抽奖活动ID
    LotteryActivityId   int64 `json:"lottery_activity_id,omitempty"`

}
