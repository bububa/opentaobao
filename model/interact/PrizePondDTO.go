package interact

// PrizePondDTO 
/* model for simplify = false
type PrizePondDTO struct {

    // 奖品列表
    
    AwardBeans  struct {
        AwardBean  []AwardBean `json:"award_bean,omitempty"`
    } `json:"award_beans,omitempty"`
    

    // 奖池开始时间
    
    StartTime   string `json:"start_time,omitempty"`
    

    // 奖池ID
    
    LotteryCode   string `json:"lottery_code,omitempty"`
    

    // 奖池结束时间
    
    EndTime   string `json:"end_time,omitempty"`
    

    // 抽奖名称
    
    LotteryName   string `json:"lottery_name,omitempty"`
    

}
*/

// PrizePondDTO 
type PrizePondDTO struct {

    // 奖品列表
    AwardBeans   []AwardBean `json:"award_beans,omitempty"`

    // 奖池开始时间
    StartTime   string `json:"start_time,omitempty"`

    // 奖池ID
    LotteryCode   string `json:"lottery_code,omitempty"`

    // 奖池结束时间
    EndTime   string `json:"end_time,omitempty"`

    // 抽奖名称
    LotteryName   string `json:"lottery_name,omitempty"`

}
