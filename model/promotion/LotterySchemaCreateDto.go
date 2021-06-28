package promotion

// LotterySchemaCreateDto 
/* model for simplify = false
type LotterySchemaCreateDto struct {

    // 关联的奖品
    
    AwardList  struct {
        LotteryAwardRelDto  []LotteryAwardRelDto `json:"lottery_award_rel_dto,omitempty"`
    } `json:"award_list,omitempty"`
    

    // 方案名称
    
    Name   string `json:"name,omitempty"`
    

    // 方案中奖概率
    
    Probability   string `json:"probability,omitempty"`
    

    // 抽奖活动id
    
    LotteryActivityId   int64 `json:"lottery_activity_id,omitempty"`
    

}
*/

// LotterySchemaCreateDto 
type LotterySchemaCreateDto struct {

    // 关联的奖品
    AwardList   []LotteryAwardRelDto `json:"award_list,omitempty"`

    // 方案名称
    Name   string `json:"name,omitempty"`

    // 方案中奖概率
    Probability   string `json:"probability,omitempty"`

    // 抽奖活动id
    LotteryActivityId   int64 `json:"lottery_activity_id,omitempty"`

}
