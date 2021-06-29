package promotion

// Feature 
type Feature struct {
    // 面额
    Amount   string `json:"amount,omitempty" xml:"amount,omitempty"`
    // 绝对生效时间开始
    EffectiveStart   string `json:"effective_start,omitempty" xml:"effective_start,omitempty"`
    // 绝对生效时间结束
    EffectiveEnd   string `json:"effective_end,omitempty" xml:"effective_end,omitempty"`
    // 生效时间类型
    EffectiveTimeMode   string `json:"effective_time_mode,omitempty" xml:"effective_time_mode,omitempty"`
    // 门槛
    StartFee   string `json:"start_fee,omitempty" xml:"start_fee,omitempty"`
    // 相对生效时间
    EffectiveInterval   string `json:"effective_interval,omitempty" xml:"effective_interval,omitempty"`
    // 相对生效时间类型
    IntervalTimeUnit   string `json:"interval_time_unit,omitempty" xml:"interval_time_unit,omitempty"`
}
