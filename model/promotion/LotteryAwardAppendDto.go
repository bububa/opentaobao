package promotion

// LotteryAwardAppendDto 
/* model for simplify = false
type LotteryAwardAppendDto struct {

    // 奖品列表
    
    AwardList  struct {
        LotteryAwardCreateDto  []LotteryAwardCreateDto `json:"lottery_award_create_dto,omitempty"`
    } `json:"award_list,omitempty"`
    

    // 抽奖活动id
    
    LotteryActivityId   int64 `json:"lottery_activity_id,omitempty"`
    

}
*/

// LotteryAwardAppendDto 
type LotteryAwardAppendDto struct {

    // 奖品列表
    AwardList   []LotteryAwardCreateDto `json:"award_list,omitempty"`

    // 抽奖活动id
    LotteryActivityId   int64 `json:"lottery_activity_id,omitempty"`

}
