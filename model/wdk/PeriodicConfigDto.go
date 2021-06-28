package wdk

// PeriodicConfigDto 
/* model for simplify = false
type PeriodicConfigDto struct {

    // 每天的什么时间阶段搞活动,精确到秒单位 例如:03:00:00_05:00:00
    
    EveryDayPeriods  struct {
        String  []string `json:"string,omitempty"`
    } `json:"every_day_periods,omitempty"`
    

    // 周期配置是否生效
    
    Periodic   bool `json:"periodic,omitempty"`
    

    // 星期几搞活动 [1:Mon;2:Tues;3:Wed;4:Thur;5:Fri;6:Sat;7:Sun]
    
    Weekdays  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"weekdays,omitempty"`
    

}
*/

// PeriodicConfigDto 
type PeriodicConfigDto struct {

    // 每天的什么时间阶段搞活动,精确到秒单位 例如:03:00:00_05:00:00
    EveryDayPeriods   []string `json:"every_day_periods,omitempty"`

    // 周期配置是否生效
    Periodic   bool `json:"periodic,omitempty"`

    // 星期几搞活动 [1:Mon;2:Tues;3:Wed;4:Thur;5:Fri;6:Sat;7:Sun]
    Weekdays   []int64 `json:"weekdays,omitempty"`

}
