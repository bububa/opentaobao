package alsc

// PointAccountOpenInfo 
type PointAccountOpenInfo struct {

    // 可用的积分总数
    
    AvailablePoint   int64 `json:"available_point,omitempty" xml:"available_point,omitempty"`
    

    // 会员ID
    
    CustomerId   string `json:"customer_id,omitempty" xml:"customer_id,omitempty"`
    

    // 冻结积分数
    
    FreezePoint   int64 `json:"freeze_point,omitempty" xml:"freeze_point,omitempty"`
    

    // 累计过期积分数
    
    TotalExpiredPoint   int64 `json:"total_expired_point,omitempty" xml:"total_expired_point,omitempty"`
    

    // 累计积分总数
    
    TotalPoint   int64 `json:"total_point,omitempty" xml:"total_point,omitempty"`
    

    // 累计可用积分数
    
    TotalUsedPoint   int64 `json:"total_used_point,omitempty" xml:"total_used_point,omitempty"`
    

    // 剩余总积分
    
    RemainPoint   int64 `json:"remain_point,omitempty" xml:"remain_point,omitempty"`
    

}
