package promotion

// LotteryAwardInstResultDto 
type LotteryAwardInstResultDto struct {

    // 奖品列表
    AwardList   []LotteryAwardDto `json:"award_list,omitempty"`

    // 奖品类型列表
    AwardTypeList   []AwardTypeDto `json:"award_type_list,omitempty"`

}
