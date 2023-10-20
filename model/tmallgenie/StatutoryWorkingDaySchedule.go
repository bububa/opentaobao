package tmallgenie

import (
	"sync"
)

// StatutoryWorkingDaySchedule 结构体
type StatutoryWorkingDaySchedule struct {
	// 响起时间（时分秒）
	Time string `json:"time,omitempty" xml:"time,omitempty"`
}

var poolStatutoryWorkingDaySchedule = sync.Pool{
	New: func() any {
		return new(StatutoryWorkingDaySchedule)
	},
}

// GetStatutoryWorkingDaySchedule() 从对象池中获取StatutoryWorkingDaySchedule
func GetStatutoryWorkingDaySchedule() *StatutoryWorkingDaySchedule {
	return poolStatutoryWorkingDaySchedule.Get().(*StatutoryWorkingDaySchedule)
}

// ReleaseStatutoryWorkingDaySchedule 释放StatutoryWorkingDaySchedule
func ReleaseStatutoryWorkingDaySchedule(v *StatutoryWorkingDaySchedule) {
	v.Time = ""
	poolStatutoryWorkingDaySchedule.Put(v)
}
