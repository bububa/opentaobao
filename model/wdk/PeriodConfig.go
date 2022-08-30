package wdk

// PeriodConfig 结构体
type PeriodConfig struct {
	// 每天的那些时间段生效
	EveryDayPeriods []string `json:"every_day_periods,omitempty" xml:"every_day_periods>string,omitempty"`
	// 一周的哪几天生效
	Weekdays []string `json:"weekdays,omitempty" xml:"weekdays>string,omitempty"`
}
