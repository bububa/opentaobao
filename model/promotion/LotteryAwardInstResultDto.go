package promotion

// LotteryAwardInstResultDto 
/* model for simplify = false
type LotteryAwardInstResultDto struct {

    // 奖品列表
    
    AwardList  struct {
        LotteryAwardDto  []LotteryAwardDto `json:"lottery_award_dto,omitempty"`
    } `json:"award_list,omitempty"`
    

    // 奖品类型列表
    
    AwardTypeList  struct {
        AwardTypeDto  []AwardTypeDto `json:"award_type_dto,omitempty"`
    } `json:"award_type_list,omitempty"`
    

}
*/

// LotteryAwardInstResultDto 
type LotteryAwardInstResultDto struct {

    // 奖品列表
    AwardList   []LotteryAwardDto `json:"award_list,omitempty"`

    // 奖品类型列表
    AwardTypeList   []AwardTypeDto `json:"award_type_list,omitempty"`

}
