package dt

// TaskDataImportSeqQry 结构体
type TaskDataImportSeqQry struct {
	// 操作人
	Operator string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 任务id
	TaskId int64 `json:"task_id,omitempty" xml:"task_id,omitempty"`
}
