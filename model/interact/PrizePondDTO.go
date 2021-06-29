package interact

// PrizePondDTO 
type PrizePondDTO struct {
    // 奖品列表
    AwardBeans   []AwardBean `json:"award_beans,omitempty" xml:"award_beans>award_bean,omitempty"`
    // 奖池开始时间
    StartTime   string `json:"start_time,omitempty" xml:"start_time,omitempty"`
    // 奖池ID
    LotteryCode   string `json:"lottery_code,omitempty" xml:"lottery_code,omitempty"`
    // 奖池结束时间
    EndTime   string `json:"end_time,omitempty" xml:"end_time,omitempty"`
    // 抽奖名称
    LotteryName   string `json:"lottery_name,omitempty" xml:"lottery_name,omitempty"`
}
