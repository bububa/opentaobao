package promotion

// LotteryAwardRelDTO 
type LotteryAwardRelDTO struct {
    // 奖品id
    AwardId   int64 `json:"award_id,omitempty" xml:"award_id,omitempty"`
    // 中奖概率
    Probability   string `json:"probability,omitempty" xml:"probability,omitempty"`
}
