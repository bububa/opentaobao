package simba

import (
	"sync"
)

// TopReportDownloadVo 结构体
type TopReportDownloadVo struct {
	// 任务id
	TaskId int64 `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

var poolTopReportDownloadVo = sync.Pool{
	New: func() any {
		return new(TopReportDownloadVo)
	},
}

// GetTopReportDownloadVo() 从对象池中获取TopReportDownloadVo
func GetTopReportDownloadVo() *TopReportDownloadVo {
	return poolTopReportDownloadVo.Get().(*TopReportDownloadVo)
}

// ReleaseTopReportDownloadVo 释放TopReportDownloadVo
func ReleaseTopReportDownloadVo(v *TopReportDownloadVo) {
	v.TaskId = 0
	poolTopReportDownloadVo.Put(v)
}
