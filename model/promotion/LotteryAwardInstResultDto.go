package promotion

// LotteryAwardInstResultDTO 
type LotteryAwardInstResultDTO struct {
    // 奖品列表
    AwardList   []LotteryAwardDTO `json:"award_list,omitempty" xml:"award_list>lottery_award_dto,omitempty"`
    // 奖品类型列表
    AwardTypeList   []AwardTypeDTO `json:"award_type_list,omitempty" xml:"award_type_list>award_type_dto,omitempty"`
}
