package promotion

// StrategyIssueResultDto 
type StrategyIssueResultDto struct {

    // 展示面额单位
    
    DisplayAmountUnit   string `json:"display_amount_unit,omitempty" xml:"display_amount_unit,omitempty"`
    

    // 扩展字段
    
    ExtraData   string `json:"extra_data,omitempty" xml:"extra_data,omitempty"`
    

    // 发放时间
    
    IssueTime   string `json:"issue_time,omitempty" xml:"issue_time,omitempty"`
    

    // 权益类型
    
    BenefitType   string `json:"benefit_type,omitempty" xml:"benefit_type,omitempty"`
    

    // 生效开始时间
    
    EffectiveStart   string `json:"effective_start,omitempty" xml:"effective_start,omitempty"`
    

    // 生效结束时间戳
    
    EffectiveEndTimestamp   int64 `json:"effective_end_timestamp,omitempty" xml:"effective_end_timestamp,omitempty"`
    

    // 外部发放实例id
    
    OuterInstanceId   string `json:"outer_instance_id,omitempty" xml:"outer_instance_id,omitempty"`
    

    // 权益code
    
    BenefitCode   string `json:"benefit_code,omitempty" xml:"benefit_code,omitempty"`
    

    // 生效结束时间
    
    EffectiveEnd   string `json:"effective_end,omitempty" xml:"effective_end,omitempty"`
    

    // 中奖记录id
    
    RecordId   int64 `json:"record_id,omitempty" xml:"record_id,omitempty"`
    

    // 用户昵称
    
    UserNick   string `json:"user_nick,omitempty" xml:"user_nick,omitempty"`
    

    // 生效开始时间戳
    
    EffectiveStartTimestamp   int64 `json:"effective_start_timestamp,omitempty" xml:"effective_start_timestamp,omitempty"`
    

    // 展示门槛
    
    DisplayStartFee   string `json:"display_start_fee,omitempty" xml:"display_start_fee,omitempty"`
    

    // 面额
    
    Amount   int64 `json:"amount,omitempty" xml:"amount,omitempty"`
    

    // 面额单位
    
    AmountUnit   string `json:"amount_unit,omitempty" xml:"amount_unit,omitempty"`
    

    // 追踪数据
    
    TrackingData   string `json:"tracking_data,omitempty" xml:"tracking_data,omitempty"`
    

    // 展示面额
    
    DisplayAmount   string `json:"display_amount,omitempty" xml:"display_amount,omitempty"`
    

    // 生效时间类型
    
    EffectiveTimeMode   string `json:"effective_time_mode,omitempty" xml:"effective_time_mode,omitempty"`
    

    // 权益标题
    
    BenefitTitle   string `json:"benefit_title,omitempty" xml:"benefit_title,omitempty"`
    

    // 权益类型名称
    
    BenefitTypeName   string `json:"benefit_type_name,omitempty" xml:"benefit_type_name,omitempty"`
    

    // 门槛
    
    StartFee   int64 `json:"start_fee,omitempty" xml:"start_fee,omitempty"`
    

    // 用户类型
    
    UserType   string `json:"user_type,omitempty" xml:"user_type,omitempty"`
    

    // 相对生效时间长度
    
    EffectiveInterval   int64 `json:"effective_interval,omitempty" xml:"effective_interval,omitempty"`
    

    // 相对生效时间单位
    
    IntervalTimeUnit   string `json:"interval_time_unit,omitempty" xml:"interval_time_unit,omitempty"`
    

}
