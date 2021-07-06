package promotion

// LotteryAwardCreateDto 结构体
type LotteryAwardCreateDto struct {
	// 奖品实例ID
	AwardInstId string `json:"award_inst_id,omitempty" xml:"award_inst_id,omitempty"`
	// 奖品类型
	AwardType int64 `json:"award_type,omitempty" xml:"award_type,omitempty"`
}
