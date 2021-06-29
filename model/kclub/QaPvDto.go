package kclub

// QaPvDTO 
type QaPvDTO struct {
    // 十天访问量
    TenDayPv   int64 `json:"ten_day_pv,omitempty" xml:"ten_day_pv,omitempty"`
    // 7天访问量
    SevenDayPv   int64 `json:"seven_day_pv,omitempty" xml:"seven_day_pv,omitempty"`
    // 5天访问量
    FiveDayPv   int64 `json:"five_day_pv,omitempty" xml:"five_day_pv,omitempty"`
    // 3天访问量
    ThreeDayPv   int64 `json:"three_day_pv,omitempty" xml:"three_day_pv,omitempty"`
    // 1天访问量
    OneDayPv   int64 `json:"one_day_pv,omitempty" xml:"one_day_pv,omitempty"`
}
