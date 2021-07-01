package promotion

// LotteryAwardInstResultDto 结构体
type LotteryAwardInstResultDto struct {
	// 奖品列表
	AwardList []LotteryAwardDto `json:"award_list,omitempty" xml:"award_list>lottery_award_dto,omitempty"`
	// 奖品类型列表
	AwardTypeList []AwardTypeDto `json:"award_type_list,omitempty" xml:"award_type_list>award_type_dto,omitempty"`
}
