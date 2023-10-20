package nlife

import (
	"sync"
)

// LogisticsLog 结构体
type LogisticsLog struct {
	// time
	Time string `json:"time,omitempty" xml:"time,omitempty"`
	// 内容
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
}

var poolLogisticsLog = sync.Pool{
	New: func() any {
		return new(LogisticsLog)
	},
}

// GetLogisticsLog() 从对象池中获取LogisticsLog
func GetLogisticsLog() *LogisticsLog {
	return poolLogisticsLog.Get().(*LogisticsLog)
}

// ReleaseLogisticsLog 释放LogisticsLog
func ReleaseLogisticsLog(v *LogisticsLog) {
	v.Time = ""
	v.Desc = ""
	poolLogisticsLog.Put(v)
}
