package promotion

// LotteryActivityRelDto 
type LotteryActivityRelDto struct {

    // 外部业务活动id
    RelationId   string `json:"relation_id,omitempty"`

    // 抽奖活动id
    LotteryAcitivityId   int64 `json:"lottery_acitivity_id,omitempty"`

}
