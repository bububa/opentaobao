package promotion

// LotteryActivityCreateDto 结构体
type LotteryActivityCreateDto struct {
	// 奖品列表
	AwardList []LotteryAwardCreateDto `json:"award_list,omitempty" xml:"award_list>lottery_award_create_dto,omitempty"`
	// 活动开始时间
	StartDate string `json:"start_date,omitempty" xml:"start_date,omitempty"`
	// 活动名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 活动结束时间
	EndDate string `json:"end_date,omitempty" xml:"end_date,omitempty"`
}
