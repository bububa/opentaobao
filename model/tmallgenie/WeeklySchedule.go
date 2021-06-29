package tmallgenie

// WeeklySchedule 
type WeeklySchedule struct {
    // 周几循环
    DaysOfWeek   []int64 `json:"days_of_week,omitempty" xml:"days_of_week>int64,omitempty"`
    // 响起时间（时分秒）
    Time   string `json:"time,omitempty" xml:"time,omitempty"`
}
