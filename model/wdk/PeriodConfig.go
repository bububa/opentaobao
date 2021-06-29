package wdk

// PeriodConfig 
type PeriodConfig struct {
    // 每天的那些时间段生效
    EveryDayPeriods   []string `json:"every_day_periods,omitempty" xml:"every_day_periods>string,omitempty"`
    // 一周的哪几天生效
    Weekdays   []int64 `json:"weekdays,omitempty" xml:"weekdays>int64,omitempty"`
}
