package wdk

// PeriodicConfigDTO 
type PeriodicConfigDTO struct {
    // 每天的什么时间阶段搞活动,精确到秒单位 例如:03:00:00_05:00:00
    EveryDayPeriods   []string `json:"every_day_periods,omitempty" xml:"every_day_periods>string,omitempty"`
    // 周期配置是否生效
    Periodic   bool `json:"periodic,omitempty" xml:"periodic,omitempty"`
    // 星期几搞活动 [1:Mon;2:Tues;3:Wed;4:Thur;5:Fri;6:Sat;7:Sun]
    Weekdays   []int64 `json:"weekdays,omitempty" xml:"weekdays>int64,omitempty"`
}
