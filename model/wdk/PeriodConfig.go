package wdk

// PeriodConfig 
/* model for simplify = false
type PeriodConfig struct {

    // 每天的那些时间段生效
    
    EveryDayPeriods  struct {
        String  []string `json:"string,omitempty"`
    } `json:"every_day_periods,omitempty"`
    

    // 一周的哪几天生效
    
    Weekdays  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"weekdays,omitempty"`
    

}
*/

// PeriodConfig 
type PeriodConfig struct {

    // 每天的那些时间段生效
    EveryDayPeriods   []string `json:"every_day_periods,omitempty"`

    // 一周的哪几天生效
    Weekdays   []int64 `json:"weekdays,omitempty"`

}
