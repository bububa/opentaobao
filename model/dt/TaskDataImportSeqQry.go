package dt

import (
	"sync"
)

// TaskDataImportSeqQry 结构体
type TaskDataImportSeqQry struct {
	// 操作人
	Operator string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 任务id
	TaskId int64 `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

var poolTaskDataImportSeqQry = sync.Pool{
	New: func() any {
		return new(TaskDataImportSeqQry)
	},
}

// GetTaskDataImportSeqQry() 从对象池中获取TaskDataImportSeqQry
func GetTaskDataImportSeqQry() *TaskDataImportSeqQry {
	return poolTaskDataImportSeqQry.Get().(*TaskDataImportSeqQry)
}

// ReleaseTaskDataImportSeqQry 释放TaskDataImportSeqQry
func ReleaseTaskDataImportSeqQry(v *TaskDataImportSeqQry) {
	v.Operator = ""
	v.TaskId = 0
	poolTaskDataImportSeqQry.Put(v)
}
