package promotion

// LotteryActivityCreateDto 
/* model for simplify = false
type LotteryActivityCreateDto struct {

    // 活动开始时间
    
    StartDate   string `json:"start_date,omitempty"`
    

    // 奖品列表
    
    AwardList  struct {
        LotteryAwardCreateDto  []LotteryAwardCreateDto `json:"lottery_award_create_dto,omitempty"`
    } `json:"award_list,omitempty"`
    

    // 活动名称
    
    Name   string `json:"name,omitempty"`
    

    // 活动结束时间
    
    EndDate   string `json:"end_date,omitempty"`
    

}
*/

// LotteryActivityCreateDto 
type LotteryActivityCreateDto struct {

    // 活动开始时间
    StartDate   string `json:"start_date,omitempty"`

    // 奖品列表
    AwardList   []LotteryAwardCreateDto `json:"award_list,omitempty"`

    // 活动名称
    Name   string `json:"name,omitempty"`

    // 活动结束时间
    EndDate   string `json:"end_date,omitempty"`

}
