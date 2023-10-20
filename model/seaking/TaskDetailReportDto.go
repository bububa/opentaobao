package seaking

import (
	"sync"
)

// TaskDetailReportDto 结构体
type TaskDetailReportDto struct {
	// 商品所在平台（ae/lazada)）
	Platform string `json:"platform,omitempty" xml:"platform,omitempty"`
	// 子任务编号，不填则表示更新所有子任务
	Idx int64 `json:"idx,omitempty" xml:"idx,omitempty"`
	// 商品id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 任务id
	TaskId int64 `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

var poolTaskDetailReportDto = sync.Pool{
	New: func() any {
		return new(TaskDetailReportDto)
	},
}

// GetTaskDetailReportDto() 从对象池中获取TaskDetailReportDto
func GetTaskDetailReportDto() *TaskDetailReportDto {
	return poolTaskDetailReportDto.Get().(*TaskDetailReportDto)
}

// ReleaseTaskDetailReportDto 释放TaskDetailReportDto
func ReleaseTaskDetailReportDto(v *TaskDetailReportDto) {
	v.Platform = ""
	v.Idx = 0
	v.ProductId = 0
	v.TaskId = 0
	poolTaskDetailReportDto.Put(v)
}
