package alihealthcrm

// AlibabaAlihealthPregnancyBaseinfoGetModel 
type AlibabaAlihealthPregnancyBaseinfoGetModel struct {

    // 上次月经的第一天,时间戳
    
    Lastday   int64 `json:"lastday,omitempty" xml:"lastday,omitempty"`
    

    // 黄体期天数
    
    LutealPhase   int64 `json:"luteal_phase,omitempty" xml:"luteal_phase,omitempty"`
    

    // 月经天数
    
    Blood   int64 `json:"blood,omitempty" xml:"blood,omitempty"`
    

    // 最长月经周期
    
    Longest   int64 `json:"longest,omitempty" xml:"longest,omitempty"`
    

    // 最短月经周期
    
    Shortest   int64 `json:"shortest,omitempty" xml:"shortest,omitempty"`
    

    // 月经周期
    
    Cycle   int64 `json:"cycle,omitempty" xml:"cycle,omitempty"`
    

    // 0：1个月以下 1：3个月以下 2半年以下 3半年以上
    
    PregnancyPeriod   int64 `json:"pregnancy_period,omitempty" xml:"pregnancy_period,omitempty"`
    

    // 月经情况 1:表示规律，0:表示不规律
    
    Regular   int64 `json:"regular,omitempty" xml:"regular,omitempty"`
    

    // 是否难孕 0:否 1:是
    
    HardPregnancy   int64 `json:"hard_pregnancy,omitempty" xml:"hard_pregnancy,omitempty"`
    

    // 备孕怀孕 0:备孕 1:怀孕
    
    Pregnancy   int64 `json:"pregnancy,omitempty" xml:"pregnancy,omitempty"`
    

    // 用户id
    
    UserId   int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
    

}
