package mos

// AlibabaMjMoscarnivalReceiveencryptData 
/* model for simplify = false
type AlibabaMjMoscarnivalReceiveencryptData struct {

    // 权益列表
    
    RightsList  struct {
        RightsList  []RightsList `json:"rights_list,omitempty"`
    } `json:"rights_list,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

    // 奖品类型
    
    PrizeType   int64 `json:"prize_type,omitempty"`
    

    // 是否有奖品
    
    HasPrize   bool `json:"has_prize,omitempty"`
    

    // 是否新会员
    
    IsNewUser   bool `json:"is_new_user,omitempty"`
    

}
*/

// AlibabaMjMoscarnivalReceiveencryptData 
type AlibabaMjMoscarnivalReceiveencryptData struct {

    // 权益列表
    RightsList   []RightsList `json:"rights_list,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // 奖品类型
    PrizeType   int64 `json:"prize_type,omitempty"`

    // 是否有奖品
    HasPrize   bool `json:"has_prize,omitempty"`

    // 是否新会员
    IsNewUser   bool `json:"is_new_user,omitempty"`

}
