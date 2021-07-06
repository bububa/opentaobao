package promotion

// LotteryAwardRelDto 结构体
type LotteryAwardRelDto struct {
	// 中奖概率
	Probability string `json:"probability,omitempty" xml:"probability,omitempty"`
	// 奖品id
	AwardId int64 `json:"award_id,omitempty" xml:"award_id,omitempty"`
}
