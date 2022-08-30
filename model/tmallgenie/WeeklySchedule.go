package tmallgenie

// WeeklySchedule 结构体
type WeeklySchedule struct {
	// 周几循环
	DaysOfWeek []string `json:"days_of_week,omitempty" xml:"days_of_week>string,omitempty"`
	// 响起时间（时分秒）
	Time string `json:"time,omitempty" xml:"time,omitempty"`
}
