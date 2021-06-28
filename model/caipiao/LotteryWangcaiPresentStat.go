package caipiao

// LotteryWangcaiPresentStat 
/* model for simplify = false
type LotteryWangcaiPresentStat struct {

    // 日期
    
    DateId   int64 `json:"date_id,omitempty"`
    

    // 送彩方的淘宝数字ID
    
    SellerId   int64 `json:"seller_id,omitempty"`
    

    // 当日赠送用户数
    
    PresentUser   int64 `json:"present_user,omitempty"`
    

    // 当日赠送彩票的注数
    
    PresentStake   int64 `json:"present_stake,omitempty"`
    

    // 当日赠送彩票的金额
    
    PresentFee   int64 `json:"present_fee,omitempty"`
    

}
*/

// LotteryWangcaiPresentStat 
type LotteryWangcaiPresentStat struct {

    // 日期
    DateId   int64 `json:"date_id,omitempty"`

    // 送彩方的淘宝数字ID
    SellerId   int64 `json:"seller_id,omitempty"`

    // 当日赠送用户数
    PresentUser   int64 `json:"present_user,omitempty"`

    // 当日赠送彩票的注数
    PresentStake   int64 `json:"present_stake,omitempty"`

    // 当日赠送彩票的金额
    PresentFee   int64 `json:"present_fee,omitempty"`

}
