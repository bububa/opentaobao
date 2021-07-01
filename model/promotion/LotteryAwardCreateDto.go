package promotion

// LotteryAwardCreateDto 
type LotteryAwardCreateDto struct {
    // 奖品类型
    AwardType   int64 `json:"award_type,omitempty" xml:"award_type,omitempty"`
    // 奖品实例ID
    AwardInstId   string `json:"award_inst_id,omitempty" xml:"award_inst_id,omitempty"`
}
