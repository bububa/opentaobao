package tmallnr

// NrZqsPlanRespDto 
type NrZqsPlanRespDto struct {

    // 每次配送件数
    
    NumPerCycle   int64 `json:"num_per_cycle,omitempty" xml:"num_per_cycle,omitempty"`
    

    // 暂停/退款提前告知的天数
    
    PauseAheadDays   int64 `json:"pause_ahead_days,omitempty" xml:"pause_ahead_days,omitempty"`
    

    // 配送时间范围，起送时间，只取时分，HH:mm格式
    
    SendStartTime   string `json:"send_start_time,omitempty" xml:"send_start_time,omitempty"`
    

    // 配送时间范围，结束时间，只取时分，HH:mm格式
    
    SendEndTime   string `json:"send_end_time,omitempty" xml:"send_end_time,omitempty"`
    

    // 每次配送的周期天数（在cycleType为1时生效，其它时候为空）1表示每天送，2表示隔1天送
    
    CycleDays   int64 `json:"cycle_days,omitempty" xml:"cycle_days,omitempty"`
    

    // 配送频率类型:1-隔N天送，2-周末送，3-工作日送
    
    CycleType   int64 `json:"cycle_type,omitempty" xml:"cycle_type,omitempty"`
    

    // 退款开始时间，注意，这个时间当天如果有配送还是会配送的，第二天开始之后的配送会取消
    
    StartRefundDate   string `json:"start_refund_date,omitempty" xml:"start_refund_date,omitempty"`
    

    // 已生成的配送计划序号及配送日期
    
    PlanList   []NrZqsPlanDetailInfoDto `json:"plan_list,omitempty" xml:"plan_list,omitempty"`
    

    // pauseInfos
    
    PauseInfos   []NrZqsPauseInfoDto `json:"pause_infos,omitempty" xml:"pause_infos,omitempty"`
    

    // 每周几送，在cycle_type=4时生效，其它时候为空， 1表示周日，2表示周一...7表示周六（以周日为每周的第一天）
    
    WeekDay   int64 `json:"week_day,omitempty" xml:"week_day,omitempty"`
    

}
