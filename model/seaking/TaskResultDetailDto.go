package seaking

import (
	"sync"
)

// TaskResultDetailDto 结构体
type TaskResultDetailDto struct {
	// 图片翻译结果
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// 子任务状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 子任务编号
	Idx int64 `json:"idx,omitempty" xml:"idx,omitempty"`
}

var poolTaskResultDetailDto = sync.Pool{
	New: func() any {
		return new(TaskResultDetailDto)
	},
}

// GetTaskResultDetailDto() 从对象池中获取TaskResultDetailDto
func GetTaskResultDetailDto() *TaskResultDetailDto {
	return poolTaskResultDetailDto.Get().(*TaskResultDetailDto)
}

// ReleaseTaskResultDetailDto 释放TaskResultDetailDto
func ReleaseTaskResultDetailDto(v *TaskResultDetailDto) {
	v.Result = ""
	v.Status = ""
	v.Idx = 0
	poolTaskResultDetailDto.Put(v)
}
