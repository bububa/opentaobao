package promotion

// LotteryActivityExtendDto 
/* model for simplify = false
type LotteryActivityExtendDto struct {

    // 活动ID
    
    Id   int64 `json:"id,omitempty"`
    

    // 活动名称
    
    Name   string `json:"name,omitempty"`
    

    // 创建时间
    
    GmtCreate   string `json:"gmt_create,omitempty"`
    

    // 修改时间
    
    GmtModified   string `json:"gmt_modified,omitempty"`
    

    // 状态（1有效）
    
    Status   int64 `json:"status,omitempty"`
    

    // 活动开始时间
    
    StartDate   string `json:"start_date,omitempty"`
    

    // 活动结束时间
    
    EndDate   string `json:"end_date,omitempty"`
    

    // 抽奖CODE
    
    RaffleCode   string `json:"raffle_code,omitempty"`
    

    // 奖品列表
    
    AwardList  struct {
        LotteryAwardExtendDto  []LotteryAwardExtendDto `json:"lottery_award_extend_dto,omitempty"`
    } `json:"award_list,omitempty"`
    

    // 活动关联的规则
    
    RuleList  struct {
        ExpressionRuleDto  []ExpressionRuleDto `json:"expression_rule_dto,omitempty"`
    } `json:"rule_list,omitempty"`
    

    // 抽奖方案列表
    
    SchemaList  struct {
        LotterySchemaDto  []LotterySchemaDto `json:"lottery_schema_dto,omitempty"`
    } `json:"schema_list,omitempty"`
    

    // 外部关联列表
    
    RelationList  struct {
        String  []string `json:"string,omitempty"`
    } `json:"relation_list,omitempty"`
    

}
*/

// LotteryActivityExtendDto 
type LotteryActivityExtendDto struct {

    // 活动ID
    Id   int64 `json:"id,omitempty"`

    // 活动名称
    Name   string `json:"name,omitempty"`

    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty"`

    // 修改时间
    GmtModified   string `json:"gmt_modified,omitempty"`

    // 状态（1有效）
    Status   int64 `json:"status,omitempty"`

    // 活动开始时间
    StartDate   string `json:"start_date,omitempty"`

    // 活动结束时间
    EndDate   string `json:"end_date,omitempty"`

    // 抽奖CODE
    RaffleCode   string `json:"raffle_code,omitempty"`

    // 奖品列表
    AwardList   []LotteryAwardExtendDto `json:"award_list,omitempty"`

    // 活动关联的规则
    RuleList   []ExpressionRuleDto `json:"rule_list,omitempty"`

    // 抽奖方案列表
    SchemaList   []LotterySchemaDto `json:"schema_list,omitempty"`

    // 外部关联列表
    RelationList   []string `json:"relation_list,omitempty"`

}
