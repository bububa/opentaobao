package interact

// LotteryProxyResult 
/* model for simplify = false
type LotteryProxyResult struct {

    // 是否中奖
    
    IsWin   bool `json:"is_win,omitempty"`
    

    // 中奖时间
    
    WinningTime   string `json:"winning_time,omitempty"`
    

    // 中奖记录id
    
    WinningRecordId   int64 `json:"winning_record_id,omitempty"`
    

    // 抽奖id
    
    LotteryId   int64 `json:"lottery_id,omitempty"`
    

    // 原因
    
    Reason   string `json:"reason,omitempty"`
    

    // 奖品类型
    
    AwardType   string `json:"award_type,omitempty"`
    

    // 奖品组id
    
    GroupId   int64 `json:"group_id,omitempty"`
    

    // 奖品id
    
    AwardId   int64 `json:"award_id,omitempty"`
    

    // 昵称
    
    Nick   string `json:"nick,omitempty"`
    

    // 用户id
    
    UserId   int64 `json:"user_id,omitempty"`
    

    // 奖品名称
    
    AwardName   string `json:"award_name,omitempty"`
    

    // 奖品拓展字段
    
    Extra   string `json:"extra,omitempty"`
    

}
*/

// LotteryProxyResult 
type LotteryProxyResult struct {

    // 是否中奖
    IsWin   bool `json:"is_win,omitempty"`

    // 中奖时间
    WinningTime   string `json:"winning_time,omitempty"`

    // 中奖记录id
    WinningRecordId   int64 `json:"winning_record_id,omitempty"`

    // 抽奖id
    LotteryId   int64 `json:"lottery_id,omitempty"`

    // 原因
    Reason   string `json:"reason,omitempty"`

    // 奖品类型
    AwardType   string `json:"award_type,omitempty"`

    // 奖品组id
    GroupId   int64 `json:"group_id,omitempty"`

    // 奖品id
    AwardId   int64 `json:"award_id,omitempty"`

    // 昵称
    Nick   string `json:"nick,omitempty"`

    // 用户id
    UserId   int64 `json:"user_id,omitempty"`

    // 奖品名称
    AwardName   string `json:"award_name,omitempty"`

    // 奖品拓展字段
    Extra   string `json:"extra,omitempty"`

}
