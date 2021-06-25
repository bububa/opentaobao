package promotion

// LotteryAwardInstQueryDto 
type LotteryAwardInstQueryDto struct {

    // 奖品类型（0优惠券1支付宝红包2粮票3红包）
    AwardTypes   []Number `json:"award_types,omitempty"`

}
