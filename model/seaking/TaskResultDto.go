package seaking

// TaskResultDto 结构体
type TaskResultDto struct {
	// 子任务列表
	TaskResultDetailList []TaskResultDetailDto `json:"task_result_detail_list,omitempty" xml:"task_result_detail_list>task_result_detail_dto,omitempty"`
	// 任务状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 完成的任务数
	FinishedCount int64 `json:"finished_count,omitempty" xml:"finished_count,omitempty"`
	// 总计任务数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
}
