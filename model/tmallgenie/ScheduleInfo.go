package tmallgenie

// ScheduleInfo 结构体
type ScheduleInfo struct {
	// DayOfWeek/DayOfMonth
	DayOfXs []int64 `json:"day_of_xs,omitempty" xml:"day_of_xs>int64,omitempty"`
	// 调度间隔
	Interval int64 `json:"interval,omitempty" xml:"interval,omitempty"`
	// 调度结束时间
	EndDate string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// 调度开始时间
	StartDate string `json:"start_date,omitempty" xml:"start_date,omitempty"`
	// 调度时间点
	Time string `json:"time,omitempty" xml:"time,omitempty"`
	// 调度频率
	Frequency string `json:"frequency,omitempty" xml:"frequency,omitempty"`
	// 调度周期
	Repeat string `json:"repeat,omitempty" xml:"repeat,omitempty"`
}
