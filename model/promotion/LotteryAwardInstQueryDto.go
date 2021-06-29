package promotion

// LotteryAwardInstQueryDTO 
type LotteryAwardInstQueryDTO struct {
    // 奖品类型（0优惠券1支付宝红包2粮票3红包）
    AwardTypes   []int64 `json:"award_types,omitempty" xml:"award_types>int64,omitempty"`
}
