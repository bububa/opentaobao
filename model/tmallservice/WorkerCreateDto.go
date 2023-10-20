package tmallservice

import (
	"sync"
)

// WorkerCreateDto 结构体
type WorkerCreateDto struct {
	// 工人ID
	WorkerId int64 `json:"worker_id,omitempty" xml:"worker_id,omitempty"`
}

var poolWorkerCreateDto = sync.Pool{
	New: func() any {
		return new(WorkerCreateDto)
	},
}

// GetWorkerCreateDto() 从对象池中获取WorkerCreateDto
func GetWorkerCreateDto() *WorkerCreateDto {
	return poolWorkerCreateDto.Get().(*WorkerCreateDto)
}

// ReleaseWorkerCreateDto 释放WorkerCreateDto
func ReleaseWorkerCreateDto(v *WorkerCreateDto) {
	v.WorkerId = 0
	poolWorkerCreateDto.Put(v)
}
