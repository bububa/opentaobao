package promotion

// LotteryAwardInstQueryDto 
/* model for simplify = false
type LotteryAwardInstQueryDto struct {

    // 奖品类型（0优惠券1支付宝红包2粮票3红包）
    
    AwardTypes  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"award_types,omitempty"`
    

}
*/

// LotteryAwardInstQueryDto 
type LotteryAwardInstQueryDto struct {

    // 奖品类型（0优惠券1支付宝红包2粮票3红包）
    AwardTypes   []int64 `json:"award_types,omitempty"`

}
