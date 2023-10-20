package seaking

import (
	"sync"
)

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

var poolTaskResultDto = sync.Pool{
	New: func() any {
		return new(TaskResultDto)
	},
}

// GetTaskResultDto() 从对象池中获取TaskResultDto
func GetTaskResultDto() *TaskResultDto {
	return poolTaskResultDto.Get().(*TaskResultDto)
}

// ReleaseTaskResultDto 释放TaskResultDto
func ReleaseTaskResultDto(v *TaskResultDto) {
	v.TaskResultDetailList = v.TaskResultDetailList[:0]
	v.Status = ""
	v.FinishedCount = 0
	v.Total = 0
	poolTaskResultDto.Put(v)
}
