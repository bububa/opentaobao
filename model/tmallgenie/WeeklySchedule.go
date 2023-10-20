package tmallgenie

import (
	"sync"
)

// WeeklySchedule 结构体
type WeeklySchedule struct {
	// 周几循环
	DaysOfWeek []string `json:"days_of_week,omitempty" xml:"days_of_week>string,omitempty"`
	// 响起时间（时分秒）
	Time string `json:"time,omitempty" xml:"time,omitempty"`
}

var poolWeeklySchedule = sync.Pool{
	New: func() any {
		return new(WeeklySchedule)
	},
}

// GetWeeklySchedule() 从对象池中获取WeeklySchedule
func GetWeeklySchedule() *WeeklySchedule {
	return poolWeeklySchedule.Get().(*WeeklySchedule)
}

// ReleaseWeeklySchedule 释放WeeklySchedule
func ReleaseWeeklySchedule(v *WeeklySchedule) {
	v.DaysOfWeek = v.DaysOfWeek[:0]
	v.Time = ""
	poolWeeklySchedule.Put(v)
}
