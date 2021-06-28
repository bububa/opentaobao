package promotion

// LotteryAwardRelDto 
/* model for simplify = false
type LotteryAwardRelDto struct {

    // 奖品id
    
    AwardId   int64 `json:"award_id,omitempty"`
    

    // 中奖概率
    
    Probability   string `json:"probability,omitempty"`
    

}
*/

// LotteryAwardRelDto 
type LotteryAwardRelDto struct {

    // 奖品id
    AwardId   int64 `json:"award_id,omitempty"`

    // 中奖概率
    Probability   string `json:"probability,omitempty"`

}
