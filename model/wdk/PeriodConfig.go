package wdk

// PeriodConfig 
type PeriodConfig struct {

    // 每天的那些时间段生效
    EveryDayPeriods   []String `json:"every_day_periods,omitempty"`

    // 一周的哪几天生效
    Weekdays   []Number `json:"weekdays,omitempty"`

}
