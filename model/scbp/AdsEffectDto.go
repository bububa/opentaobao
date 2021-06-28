package scbp

// AdsEffectDto 
/* model for simplify = false
type AdsEffectDto struct {

    // 曝光
    
    Impr   int64 `json:"impr,omitempty"`
    

    // 点击
    
    Click   int64 `json:"click,omitempty"`
    

    // 消耗
    
    Cost   int64 `json:"cost,omitempty"`
    

    // 推广时长
    
    OnlineMin   int64 `json:"online_min,omitempty"`
    

}
*/

// AdsEffectDto 
type AdsEffectDto struct {

    // 曝光
    Impr   int64 `json:"impr,omitempty"`

    // 点击
    Click   int64 `json:"click,omitempty"`

    // 消耗
    Cost   int64 `json:"cost,omitempty"`

    // 推广时长
    OnlineMin   int64 `json:"online_min,omitempty"`

}
