package aesolution

import (
	"sync"
)

// BatchOperationJobDto 结构体
type BatchOperationJobDto struct {
	// The status of the feed
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// feed type
	OperationType string `json:"operation_type,omitempty" xml:"operation_type,omitempty"`
	// job id
	JobId int64 `json:"job_id,omitempty" xml:"job_id,omitempty"`
}

var poolBatchOperationJobDto = sync.Pool{
	New: func() any {
		return new(BatchOperationJobDto)
	},
}

// GetBatchOperationJobDto() 从对象池中获取BatchOperationJobDto
func GetBatchOperationJobDto() *BatchOperationJobDto {
	return poolBatchOperationJobDto.Get().(*BatchOperationJobDto)
}

// ReleaseBatchOperationJobDto 释放BatchOperationJobDto
func ReleaseBatchOperationJobDto(v *BatchOperationJobDto) {
	v.Status = ""
	v.OperationType = ""
	v.JobId = 0
	poolBatchOperationJobDto.Put(v)
}
