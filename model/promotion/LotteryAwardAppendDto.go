package promotion

// LotteryAwardAppendDTO 
type LotteryAwardAppendDTO struct {
    // 奖品列表
    AwardList   []LotteryAwardCreateDTO `json:"award_list,omitempty" xml:"award_list>lottery_award_create_dto,omitempty"`
    // 抽奖活动id
    LotteryActivityId   int64 `json:"lottery_activity_id,omitempty" xml:"lottery_activity_id,omitempty"`
}
