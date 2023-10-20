package tmallgenie

import (
	"sync"
)

// OnceSchedule 结构体
type OnceSchedule struct {
	// 响起日期和时间（年月日时分秒）
	Datetime string `json:"datetime,omitempty" xml:"datetime,omitempty"`
}

var poolOnceSchedule = sync.Pool{
	New: func() any {
		return new(OnceSchedule)
	},
}

// GetOnceSchedule() 从对象池中获取OnceSchedule
func GetOnceSchedule() *OnceSchedule {
	return poolOnceSchedule.Get().(*OnceSchedule)
}

// ReleaseOnceSchedule 释放OnceSchedule
func ReleaseOnceSchedule(v *OnceSchedule) {
	v.Datetime = ""
	poolOnceSchedule.Put(v)
}
