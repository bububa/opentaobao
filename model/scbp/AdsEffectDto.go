package scbp

// AdsEffectDto 
type AdsEffectDto struct {

    // 曝光
    
    Impr   int64 `json:"impr,omitempty" xml:"impr,omitempty"`
    

    // 点击
    
    Click   int64 `json:"click,omitempty" xml:"click,omitempty"`
    

    // 消耗
    
    Cost   int64 `json:"cost,omitempty" xml:"cost,omitempty"`
    

    // 推广时长
    
    OnlineMin   int64 `json:"online_min,omitempty" xml:"online_min,omitempty"`
    

}
